package app

import (
	"context"

	"github.com/elojah/game_03/pkg/twitch"
)

var _ twitch.App = (*App)(nil)

type App struct {
	twitch.Client

	clientID string
}

func (a *App) Dial(ctx context.Context, cfg twitch.Config) error {
	a.clientID = cfg.ClientID

	return nil
}

func (a *App) Close(ctx context.Context) error {
	a.clientID = ""

	return nil
}

func (a App) ClientID() string {
	return a.clientID
}
