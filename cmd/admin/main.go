package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	admingrpc "github.com/elojah/game_03/cmd/admin/grpc"
	migrateapp "github.com/elojah/game_03/pkg/migrate/app"
	roomapp "github.com/elojah/game_03/pkg/room/app"
	roomscylla "github.com/elojah/game_03/pkg/room/scylla"
	ggrpc "github.com/elojah/go-grpc"
	glog "github.com/elojah/go-log"
	"github.com/elojah/go-scylla"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

const (
	// Time allocated for init phase (connections + setup).
	initTO = 30 * time.Second
)

var version string

type closer interface {
	Close(context.Context) error
}

type closers []closer

func (cs closers) Close(ctx context.Context) error {
	var result *multierror.Error

	for _, c := range cs {
		result = multierror.Append(result, c.Close(ctx))
	}

	return result.ErrorOrNil()
}

// run services.
func run(prog string, filename string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, initTO)
	// no need for defer cancel
	_ = cancel

	var cs []closer

	logs := glog.Service{}
	if err := logs.Dial(ctx, glog.Config{}); err != nil {
		fmt.Println("failed to dial logger")

		return
	}

	cs = append(cs, &logs)

	log.Logger = logs.With().Str("version", version).Str("exe", prog).Logger()

	// read config
	cfg := config{}
	if err := cfg.Populate(ctx, filename); err != nil {
		log.Error().Err(err).Msg("failed to read config")

		return
	}

	// init ghttp web server
	scyllas := scylla.Service{}

	if err := scyllas.Dial(ctx, cfg.Scylla); err != nil {
		log.Error().Err(err).Msg("failed to dial scylla")

		return
	}

	cs = append(cs, &scyllas)

	migrateApp := migrateapp.App{
		Service: &scyllas,
	}

	roomStore := &roomscylla.Store{Service: scyllas}
	roomApp := roomapp.App{
		Store:          roomStore,
		StoreCell:      roomStore,
		StoreWorld:     roomStore,
		StoreUser:      roomStore,
		StoreWorldCell: roomStore,
	}

	h := handler{
		migrate: &migrateApp,
		room:    roomApp,
	}

	// init grpc api server
	grpcadmin := ggrpc.Service{}
	grpcadmin.Register = func() {
		reflection.Register(grpcadmin.Server)
		admingrpc.RegisterAdminServer(grpcadmin.Server, &h)
	}

	if err := grpcadmin.Dial(ctx, cfg.GRPC); err != nil {
		log.Error().Err(err).Msg("failed to dial grpc")

		return
	}

	cs = append(cs, &grpcadmin)

	go func() {
		if err := grpcadmin.Serve(ctx); err != nil {
			log.Error().Err(err).Msg("failed to serve grpc")
		}
	}()

	log.Info().Msg("admin up")

	// listen for signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for sig := range c {
		switch sig {
		case syscall.SIGHUP:
			fallthrough
		case syscall.SIGINT:
			fallthrough
		case syscall.SIGTERM:
			if err := closers(cs).Close(ctx); err != nil {
				fmt.Printf("error closing service: %s\n", err.Error())

				return
			}

			cancel()

			fmt.Println("successfully closed service")

			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 { //nolint: gomnd
		fmt.Printf("Usage: ./%s configfile\n", args[0])

		return
	}

	run(args[0], args[1])
}
