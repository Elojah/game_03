package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID   *ulid.ID
	IDs  []ulid.ID
	PCID *ulid.ID

	X *int64
	Y *int64

	State []byte
	Size  int
}

type Store interface {
	Insert(context.Context, E) error
	Fetch(context.Context, Filter) (E, error)
	FetchMany(context.Context, Filter) ([]E, []byte, error)
	Delete(context.Context, Filter) error
}

type App interface {
	Store
	StorePC
}
