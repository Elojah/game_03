package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterCell struct {
	WorldID *ulid.ID
	X       *int64
	Y       *int64
}

type StoreCell interface {
	InsertCells(context.Context, ...Cell) error
	FetchCell(context.Context, FilterCell) (Cell, error)
	DeleteCell(context.Context, FilterCell) error
}
