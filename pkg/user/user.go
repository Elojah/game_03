package user

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type App interface {
	Store
	StoreSession
	// CacheSession

	Auth(ctx context.Context) (Session, error)
}

type Filter struct {
	ID       *ulid.ID
	TwitchID *string
}

type Store interface {
	Insert(context.Context, U) error
	Fetch(context.Context, Filter) (U, error)
	Delete(context.Context, Filter) error
}
