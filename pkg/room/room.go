package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	// IDs      *ulid.IDs
	IDs     []ulid.ID
	OwnerID *ulid.ID

	// Pagination
	Size  int    `db:"-"`
	State []byte `db:"-"`
}

type Store interface {
	Insert(context.Context, R) error
	Fetch(context.Context, Filter) (R, error)
	FetchMany(context.Context, Filter) ([]R, []byte, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Store
	StoreWorld
	StoreCell
}
