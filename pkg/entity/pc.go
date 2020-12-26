package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterPC struct {
	ID     *ulid.ID
	UserID *ulid.ID
}

type StorePC interface {
	UpsertPC(context.Context, PC) error
	FetchPC(context.Context, FilterPC) (PC, error)
	FetchManyPC(context.Context, FilterPC, func(PC) error) error
	DeletePC(context.Context, FilterPC) error
	DeleteManyPC(context.Context, FilterPC) error
}
