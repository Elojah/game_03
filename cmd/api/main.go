package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	apigrpc "github.com/elojah/game_03/cmd/api/grpc"
	cookieapp "github.com/elojah/game_03/pkg/cookie/app"
	cookieredis "github.com/elojah/game_03/pkg/cookie/redis"
	entityapp "github.com/elojah/game_03/pkg/entity/app"
	entityredis "github.com/elojah/game_03/pkg/entity/redis"
	entityscylla "github.com/elojah/game_03/pkg/entity/scylla"
	roomapp "github.com/elojah/game_03/pkg/room/app"
	roomscylla "github.com/elojah/game_03/pkg/room/scylla"
	userapp "github.com/elojah/game_03/pkg/user/app"
	userredis "github.com/elojah/game_03/pkg/user/redis"
	userscylla "github.com/elojah/game_03/pkg/user/scylla"
	ggrpc "github.com/elojah/go-grpc"
	"github.com/elojah/go-grpcweb"
	ghttp "github.com/elojah/go-http"
	glog "github.com/elojah/go-log"
	"github.com/elojah/go-redis"
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
		if c != nil {
			result = multierror.Append(result, c.Close(ctx))
		}
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

	cookieCache := &cookieredis.Cache{Service: rediss}
	cookieApp := &cookieapp.A{
		CacheKeys: cookieCache,
	}

	// init http api server
	https := ghttp.Service{}

	if err := https.Dial(ctx, cfg.HTTP); err != nil {
		log.Error().Err(err).Msg("failed to dial http")

		return
	}

	cs = append(cs, &https)

	entityCache := &entityredis.Cache{Service: rediss}
	entityStore := &entityscylla.Store{Service: scyllas}
	entityApp := entityapp.App{
		Cache:          entityCache,
		Store:          entityStore,
		StoreAnimation: entityStore,
		StoreBackup:    entityStore,
		StorePC:        entityStore,
		StoreTemplate:  entityStore,
		StoreSpawn:     entityStore,
	}

	roomStore := &roomscylla.Store{Service: scyllas}
	roomApp := roomapp.App{
		Store:           roomStore,
		StorePublic:     roomStore,
		StoreCell:       roomStore,
		StoreWorld:      roomStore,
		StoreUser:       roomStore,
		StoreWorldCell:  roomStore,
		StoreWorldSpawn: roomStore,
	}

	userStore := &userscylla.Store{Service: scyllas}
	userCache := &userredis.Cache{Service: rediss}
	userApp := userapp.App{
		Store:        userStore,
		StoreSession: userStore,
		CacheSession: userCache,
		Cookie:       cookieApp,
	}

	if err := userApp.Dial(ctx, cfg.Session); err != nil {
		log.Error().Err(err).Msg("failed to dial user application")

		return
	}

	// init handler
	h := handler{
		entity: entityApp,
		room:   roomApp,
		user:   userApp,
	}

	// init grpc ONLY server
	grpcapi := ggrpc.Service{}
	grpcapi.Register = func() {
		reflection.Register(grpcapi.Server)
		apigrpc.RegisterAPIServer(grpcapi.Server, &h)
	}

	if err := grpcapi.Dial(ctx, cfg.GRPC); err != nil {
		log.Error().Err(err).Msg("failed to dial grpc")

		return
	}

	cs = append(cs, &grpcapi)

	go func() {
		if err := grpcapi.Serve(ctx); err != nil {
			log.Error().Err(err).Msg("failed to serve grpc")
		}
	}()

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
		if err := https.Server.ListenAndServeTLS("", ""); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info().Msg("http server closed")
			} else {
				log.Error().Err(err).Msg("failed to serve http")
			}
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
	if len(args) != 2 { //nolint: gomnd
		fmt.Printf("Usage: ./%s configfile\n", args[0])

		return
	}

	run(args[0], args[1])
}
