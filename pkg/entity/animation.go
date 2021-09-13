package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterAnimation struct {
	EntityIDs []ulid.ID

	State []byte
	Size  int
}

type StoreAnimation interface {
	InsertAnimation(context.Context, Animation) error
	FetchAnimation(context.Context, FilterAnimation) (Animation, error)
	FetchManyAnimation(context.Context, FilterAnimation) ([]Animation, []byte, error)
	DeleteAnimation(context.Context, FilterAnimation) error
}
