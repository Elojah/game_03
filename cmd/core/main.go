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

	coregrpc "github.com/elojah/game_03/cmd/core/grpc"
	abilityapp "github.com/elojah/game_03/pkg/ability/app"
	abilityredis "github.com/elojah/game_03/pkg/ability/redis"
	abilityscylla "github.com/elojah/game_03/pkg/ability/scylla"
	cookieapp "github.com/elojah/game_03/pkg/cookie/app"
	cookieredis "github.com/elojah/game_03/pkg/cookie/redis"
	entityapp "github.com/elojah/game_03/pkg/entity/app"
	entityredis "github.com/elojah/game_03/pkg/entity/redis"
	entityscylla "github.com/elojah/game_03/pkg/entity/scylla"
	eventapp "github.com/elojah/game_03/pkg/event/app"
	eventredis "github.com/elojah/game_03/pkg/event/redis"
	roomapp "github.com/elojah/game_03/pkg/room/app"
	roomscylla "github.com/elojah/game_03/pkg/room/scylla"
	rtcapp "github.com/elojah/game_03/pkg/rtc/app"
	rtcmem "github.com/elojah/game_03/pkg/rtc/mem"
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

	// init http core server
	https := ghttp.Service{}

	if err := https.Dial(ctx, cfg.HTTP); err != nil {
		log.Error().Err(err).Msg("failed to dial http")

		return
	}

	cs = append(cs, &https)

	abilityCache := &abilityredis.Cache{Service: rediss}
	abilityStore := &abilityscylla.Store{Service: scyllas}
	abilityApp := abilityapp.App{
		Cache: abilityCache,
		Store: abilityStore,
	}

	entityCache := &entityredis.Cache{Service: rediss}
	entityStore := &entityscylla.Store{Service: scyllas}
	entityApp := entityapp.App{
		Cache:              entityCache,
		Store:              entityStore,
		StoreAbility:       entityStore,
		CacheAbility:       entityCache,
		StoreAnimation:     entityStore,
		StoreBackup:        entityStore,
		StorePC:            entityStore,
		StorePCPreferences: entityStore,
		StoreTemplate:      entityStore,
		StoreSpawn:         entityStore,
		Ability:            abilityApp,
	}

	roomStore := &roomscylla.Store{Service: scyllas}
	roomApp := roomapp.App{
		Store:              roomStore,
		StorePublic:        roomStore,
		StoreCell:          roomStore,
		StoreWorld:         roomStore,
		StoreUser:          roomStore,
		StoreWorldCell:     roomStore,
		StoreWorldSpawn:    roomStore,
		StoreWorldWaypoint: roomStore,
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

	rtcStore := rtcmem.NewStore()
	rtcApp := rtcapp.App{
		Store: &rtcStore,
	}

	eventCache := &eventredis.Cache{Service: rediss}
	eventApp := eventapp.App{
		Cache:   eventCache,
		CacheQ:  eventCache,
		Entity:  entityApp,
		Ability: abilityApp,
		Room:    roomApp,
	}

	// init handler
	h := handler{
		ability: abilityApp,
		entity:  entityApp,
		event:   eventApp,
		room:    roomApp,
		rtc:     rtcApp,
		user:    userApp,
	}

	if err := h.Dial(ctx); err != nil {
		log.Error().Err(err).Msg("failed to dial handler")

		return
	}

	// init grpc ONLY server
	grpccore := ggrpc.Service{}
	grpccore.Register = func() {
		reflection.Register(grpccore.Server)
		coregrpc.RegisterCoreServer(grpccore.Server, &h)
	}

	if err := grpccore.Dial(ctx, cfg.GRPC); err != nil {
		log.Error().Err(err).Msg("failed to dial grpc")

		return
	}

	cs = append(cs, &grpccore)

	go func() {
		if err := grpccore.Serve(ctx); err != nil {
			log.Error().Err(err).Msg("failed to serve grpc")
		}
	}()

	// init grpc core server
	grpcwcore := grpcweb.Service{}
	grpcwcore.Register = func() {
		coregrpc.RegisterCoreServer(grpcwcore.Server, &h)
		// reflection.Register(grpcwcore.Server)
		https.Handler = grpcwcore.WrappedGrpcServer
	}

	if err := grpcwcore.Dial(ctx, cfg.GRPCWeb); err != nil {
		log.Error().Err(err).Msg("failed to dial grpcweb")

		return
	}

	cs = append(cs, &grpcwcore)

	// serve grpcweb core
	go func() {
		if err := https.Server.ListenAndServeTLS("", ""); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info().Msg("http server closed")
			} else {
				log.Error().Err(err).Msg("failed to serve http")
			}
		}
	}()

	log.Info().Msg("core up")

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
