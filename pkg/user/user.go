package user

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type App interface {
	Store
	StoreSession

	Auth(ctx context.Context) (Session, error)
}

type Filter struct {
	ID       *ulid.ID
	IDs      []ulid.ID
	TwitchID *string

	State []byte
	Size  int
}

type Store interface {
	Insert(context.Context, U) error
	Fetch(context.Context, Filter) (U, error)
	FetchMany(context.Context, Filter) ([]U, []byte, error)
	Delete(context.Context, Filter) error
}
