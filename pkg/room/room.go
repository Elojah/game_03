package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Filter struct {
	ID *ulid.ID
}

type Store interface {
	Insert(context.Context, R) error
	Fetch(context.Context, Filter) (R, error)
	Delete(context.Context, Filter) error
}
