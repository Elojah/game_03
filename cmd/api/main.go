package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	apigrpc "github.com/elojah/game_03/cmd/api/grpc"
	"github.com/elojah/go-firebase"
	"github.com/elojah/go-grpcweb"
	ghttp "github.com/elojah/go-http"
	glog "github.com/elojah/go-log"
	"github.com/elojah/go-redis"
	"github.com/elojah/go-scylla"
	"github.com/hashicorp/go-multierror"
	"github.com/rs/zerolog/log"
	_ "google.golang.org/grpc/encoding/gzip"
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

	// init Scylla storage
	scyllas := scylla.Service{}
	if err := scyllas.Dial(ctx, cfg.Scylla); err != nil {
		log.Error().Err(err).Msg("failed to dial scylla")

		return
	}

	cs = append(cs, &scyllas)

	// init redis storage
	rediss := redis.Service{}
	if err := rediss.Dial(ctx, cfg.Redis); err != nil {
		log.Error().Err(err).Msg("failed to dial redis")

		return
	}

	cs = append(cs, &rediss)

	// init firebase storage
	fbs := firebase.Service{}
	if err := fbs.Dial(ctx, cfg.Firebase); err != nil {
		log.Error().Err(err).Msg("failed to dial firebase")

		return
	}

	cs = append(cs, &fbs)

	// init http api server
	https := ghttp.Service{}

	if err := https.Dial(ctx, cfg.HTTP); err != nil {
		log.Error().Err(err).Msg("failed to dial http")

		return
	}

	cs = append(cs, &https)

	// init handler
	h := handler{}

	// init grpc api server
	grpcwapi := grpcweb.Service{}
	grpcwapi.Register = func() {
		apigrpc.RegisterAPIServer(grpcwapi.Server, &h)
		// reflection.Register(grpcwapi.Server)
		https.Handler = grpcwapi.WrappedGrpcServer
	}

	if err := grpcwapi.Dial(ctx, cfg.GRPCWeb); err != nil {
		log.Error().Err(err).Msg("failed to dial grpcweb")

		return
	}

	cs = append(cs, &grpcwapi)

	// serve grpcweb api
	go func() {
		if err := https.Server.ListenAndServe(); err != nil {
			log.Error().Err(err).Msg("failed to serve http")
		}
	}()
	log.Info().Msg("api up")

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
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, initTO)

			defer cancel()

			if err := closers(cs).Close(ctx); err != nil {
				fmt.Printf("error closing service: %s\n", err.Error())

				return
			}

			fmt.Println("successfully closed service")

			return
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 { // nolint: gomnd
		fmt.Printf("Usage: ./%s configfile\n", args[0])

		return
	}

	run(args[0], args[1])
}
