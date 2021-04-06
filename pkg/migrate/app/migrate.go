package app

import (
	"context"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/go-scylla"
	"github.com/scylladb/gocqlx/migrate"
)

type App struct {
	*scylla.Service
}

func (a *App) Up(ctx context.Context, dir string) error {
	return migrate.Migrate(ctx, a.Session.Session, dir)
}

func (a *App) Down(context.Context, string) error {
	return errors.ErrNotImplemented{Version: "soontm"}
}
