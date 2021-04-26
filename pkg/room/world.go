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

func (w World) NewCells() []Cell {
	nh := w.Height / w.CellHeight
	nw := w.Width / w.CellWidth

	cells := make([]Cell, nh*nw)
	// for
	return cells
}
