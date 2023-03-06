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

	authgrpc "github.com/elojah/game_03/cmd/auth/grpc"
	cookieapp "github.com/elojah/game_03/pkg/cookie/app"
	cookieredis "github.com/elojah/game_03/pkg/cookie/redis"
	ggrpc "github.com/elojah/go-grpc"
	ghttp "github.com/elojah/go-http"
	glog "github.com/elojah/go-log"
	"github.com/elojah/go-redis"
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

	// init http web server
	https := ghttp.Service{}

	if err := https.Dial(ctx, cfg.HTTP); err != nil {
		log.Error().Err(err).Msg("failed to dial http")

		return
	}

	cs = append(cs, &https)

	// init redis storage
	rediss := redis.Service{}
	if err := rediss.Dial(ctx, cfg.Redis); err != nil {
		log.Error().Err(err).Msg("failed to dial redis")

		return
	}

	cs = append(cs, &rediss)

	cookieStore := &cookieredis.Store{Service: rediss}
	cookieApp := &cookieapp.A{
		StoreKeys: cookieStore,
	}

	// sync local keys every 60 seconds
	go func() {
		if err := cookieApp.AutoSyncKeys(context.Background(), 60); err != nil { //nolint: gomnd
			log.Error().Err(err).Msg("failed to auto sync keys")
		}
	}()

	authclient := ggrpc.Client{}
	if err := authclient.Dial(ctx, cfg.AuthClient); err != nil {
		log.Error().Err(err).Msg("failed to dial auth")

		return
	}

	cs = append(cs, &authclient)

	h := handler{
		cookie:     cookieApp,
		AuthClient: authgrpc.NewAuthClient(authclient.ClientConn),
	}

	// auth
	https.Router.Path("/signin_google").HandlerFunc(h.signinGoogle)
	https.Router.Path("/signin_twitch").HandlerFunc(h.signinTwitch)
	https.Router.Path("/refresh_token").HandlerFunc(h.refreshToken)
	// serve static dir
	https.Router.PathPrefix("/").Handler(http.FileServer(http.Dir(cfg.Web.Static)))

	// serve http web
	go func() {
		if err := https.Server.ListenAndServeTLS("", ""); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info().Msg("http server closed")
			} else {
				log.Error().Err(err).Msg("failed to serve http")
			}
		}
	}()
	log.Info().Msg("web dashboard up")

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
