package app

import (
	"context"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/google"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

var _ google.App = (*App)(nil)

type App struct {
	config oauth2.Config
}

func (a *App) Dial(ctx context.Context, cfg oauth2.Config) error {
	a.config = cfg

	return nil
}

func (a *App) Close(ctx context.Context) error {
	a.config = oauth2.Config{}

	return nil
}

func (a App) OAuth() oauth2.Config {
	return a.config
}

func (a App) Signin(ctx context.Context, token string) (string, error) {
	// #Validate token
	p, err := idtoken.Validate(ctx, token, a.config.ClientID)
	if err != nil {
		return "", err
	}

	// #Retrieve Google ID
	subI, ok := p.Claims["sub"]
	if !ok {
		err := errors.ErrInvalidClaim{Claim: "not found"}
		return "", err
	}

	sub, ok := subI.(string)
	if !ok {
		err := errors.ErrInvalidClaim{Claim: "invalid format"}
		return "", err
	}

	return sub, nil
}
