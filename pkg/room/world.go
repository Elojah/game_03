package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterWorld struct {
	ID *ulid.ID
}

type StoreWorld interface {
	InsertWorld(context.Context, World) error
	FetchWorld(context.Context, FilterWorld) (World, error)
	DeleteWorld(context.Context, FilterWorld) error
}

func (w World) Cell(x int64, y int64) (int64, int64) {
	return w.Width / w.CellWidth, w.Height / w.CellHeight
}
