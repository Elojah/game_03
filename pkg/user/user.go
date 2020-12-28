package user

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type App interface {
	Store
	CacheSession

	FetchWithUpsert(context.Context, string) (U, error)
}

type Filter struct {
	ID       *ulid.ID
	TwitchID *string
}

type Store interface {
	Upsert(context.Context, U) error
	Fetch(context.Context, Filter) (U, error)
}
