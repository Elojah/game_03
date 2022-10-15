package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	authgrpc "github.com/elojah/game_03/cmd/auth/grpc"
	cookieapp "github.com/elojah/game_03/pkg/cookie/app"
	cookieredis "github.com/elojah/game_03/pkg/cookie/redis"
	googleapp "github.com/elojah/game_03/pkg/google/app"
	twitchapp "github.com/elojah/game_03/pkg/twitch/app"
	twitchhttp "github.com/elojah/game_03/pkg/twitch/http"
	userapp "github.com/elojah/game_03/pkg/user/app"
	userscylla "github.com/elojah/game_03/pkg/user/scylla"
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

	// init redis storage
	rediss := redis.Service{}
	if err := rediss.Dial(ctx, cfg.Redis); err != nil {
		log.Error().Err(err).Msg("failed to dial redis")

		return
	}

	cs = append(cs, &rediss)

	// init domain
	twitchApp := twitchapp.App{
		Client: &twitchhttp.Client{
			Client: http.DefaultClient,
		},
	}
	if err := twitchApp.Dial(ctx, cfg.Twitch); err != nil {
		log.Error().Err(err).Msg("failed to dial twitch app")

		return
	}

	googleApp := googleapp.App{}
	if err := googleApp.Dial(ctx, cfg.Google); err != nil {
		log.Error().Err(err).Msg("failed to dial google app")

		return
	}

	cs = append(cs, &googleApp)

	cookieStore := &cookieredis.Store{Service: rediss}
	cookieApp := &cookieapp.A{
		StoreKeys: cookieStore,
	}

	userStore := &userscylla.Store{Service: scyllas}
	userApp := userapp.App{
		Store:        userStore,
		StoreSession: userStore,
		Cookie:       cookieApp,
	}

	// init handler
	h := handler{
		twitch: twitchApp,
		google: googleApp,

		user: userApp,
	}

	// init grpc api server
	grpcauth := ggrpc.Service{}
	grpcauth.Register = func() {
		reflection.Register(grpcauth.Server)
		authgrpc.RegisterAuthServer(grpcauth.Server, &h)
	}

	if err := grpcauth.Dial(ctx, cfg.GRPC); err != nil {
		log.Error().Err(err).Msg("failed to dial grpc")

		return
	}

	cs = append(cs, &grpcauth)

	go func() {
		if err := grpcauth.Serve(ctx); err != nil {
			log.Error().Err(err).Msg("failed to serve grpc")
		}
	}()

	log.Info().Msg("auth up")

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
