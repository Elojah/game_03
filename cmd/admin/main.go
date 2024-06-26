package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	admingrpc "github.com/elojah/game_03/cmd/admin/grpc"
	abilityapp "github.com/elojah/game_03/pkg/ability/app"
	abilityredis "github.com/elojah/game_03/pkg/ability/redis"
	abilityscylla "github.com/elojah/game_03/pkg/ability/scylla"
	"github.com/elojah/game_03/pkg/cookie"
	cookieapp "github.com/elojah/game_03/pkg/cookie/app"
	cookieredis "github.com/elojah/game_03/pkg/cookie/redis"
	entityapp "github.com/elojah/game_03/pkg/entity/app"
	entityredis "github.com/elojah/game_03/pkg/entity/redis"
	entityscylla "github.com/elojah/game_03/pkg/entity/scylla"
	factionapp "github.com/elojah/game_03/pkg/faction/app"
	factionscylla "github.com/elojah/game_03/pkg/faction/scylla"
	migrateapp "github.com/elojah/game_03/pkg/migrate/app"
	roomapp "github.com/elojah/game_03/pkg/room/app"
	roomscylla "github.com/elojah/game_03/pkg/room/scylla"
	tileapp "github.com/elojah/game_03/pkg/tile/app"
	tilescylla "github.com/elojah/game_03/pkg/tile/scylla"
	ggrpc "github.com/elojah/go-grpc"
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

	// init ghttp web server
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

	// setup initial cookie keys
	if err := cookieApp.Setup(ctx, cookie.NKeys); err != nil {
		log.Error().Err(err).Msg("failed to setup cookie keys")

		return
	}

	migrateApp := &migrateapp.App{
		Service: &scyllas,
	}

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

	factionStore := &factionscylla.Store{Service: scyllas}
	factionApp := factionapp.App{
		Store:   factionStore,
		StorePC: factionStore,
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

		Entity:  entityApp,
		Faction: factionApp,
	}

	tileStore := &tilescylla.Store{Service: scyllas}
	tileApp := tileapp.App{
		StoreSet: tileStore,
	}

	h := handler{
		ability: abilityApp,
		cookie:  cookieApp,
		entity:  entityApp,
		faction: factionApp,
		migrate: migrateApp,
		room:    roomApp,
		tile:    tileApp,
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
