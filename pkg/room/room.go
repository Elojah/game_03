package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID       *ulid.ID
	IDs      []ulid.ID
	OwnerID  *ulid.ID
	OwnerIDs []ulid.ID

	State []byte
	Size  int
}

type Store interface {
	Insert(context.Context, R) error
	Fetch(context.Context, Filter) (R, error)
	FetchMany(context.Context, Filter) ([]R, []byte, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Store
	StoreCell
	StoreWorld
	StoreWorldCell
}
