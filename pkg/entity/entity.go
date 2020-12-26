package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID  *ulid.ID
	IDs []ulid.ID

	TS *uint64
}

// Store is persistent entity state storage.
type Store interface {
	Upsert(context.Context, E) error
	Fetch(context.Context, Filter) (E, error)
	FetchMany(context.Context, Filter, func(E) error) error
	Delete(context.Context, Filter) error
	DeleteMany(context.Context, Filter) error
}

// Cache is non-persistent storage for in game operations.
type Cache interface {
	Upsert(context.Context, E) error
	Fetch(context.Context, Filter) (E, error)
	FetchMany(context.Context, Filter, func(E) error) error
	Delete(context.Context, Filter) error
	DeleteMany(context.Context, Filter) error
}
