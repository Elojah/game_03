package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type KeyCell struct {
	WorldID ulid.ID
	X       int64
	Y       int64
}

type FilterCell struct {
	Keys []KeyCell

	State []byte
	Size  int
}

type StoreCell interface {
	InsertCell(context.Context, Cell) error
	FetchCell(context.Context, FilterCell) (Cell, error)
	FetchManyCell(context.Context, FilterCell) ([]Cell, []byte, error)
	DeleteCell(context.Context, FilterCell) error
}
