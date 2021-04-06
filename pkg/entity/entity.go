package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID *ulid.ID
}

type Store interface {
	Insert(context.Context, E) error
	Fetch(context.Context, Filter) (E, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Store
	StorePC
}
