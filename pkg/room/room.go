package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID      *ulid.ID
	IDs     []ulid.ID
	OwnerID *ulid.ID
}

type Store interface {
	Insert(context.Context, R) error
	Fetch(context.Context, Filter) (R, error)
	FetchMany(context.Context, Filter) ([]R, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Store
	StoreWorld
	StoreCell
}
