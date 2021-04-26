package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterWorldCell struct {
	WorldID ulid.ID
	X       int64
	Y       int64
}

type FilterCell struct {
	ID  *ulid.ID
	IDs []ulid.ID

	State []byte
	Size  int
}

type StoreCell interface {
	InsertCell(context.Context, Cell) error
	FetchCell(context.Context, FilterCell) (Cell, error)
	FetchManyCell(context.Context, FilterCell) ([]Cell, []byte, error)
	DeleteCell(context.Context, FilterCell) error
}

type StoreWorldCell interface {
	InsertWorldCell(context.Context, WorldCell) error
	FetchWorldCell(context.Context, FilterWorldCell) (WorldCell, error)
	DeleteWorldCell(context.Context, FilterWorldCell) error
}
