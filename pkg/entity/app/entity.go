package app

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
)

var _ entity.App = (*App)(nil)

type App struct {
	entity.Store
	entity.StoreAnimation
	entity.StoreBackup
	entity.StoreNPC
	entity.StorePC
	entity.CachePCConnect
}

func (a *App) FetchManyParallelByCell(ctx context.Context, f entity.Filter) ([]entity.E, []byte, error) {
	return nil, nil, nil
}
