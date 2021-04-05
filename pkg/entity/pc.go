package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterPC struct {
	ID *ulid.ID
}

type StorePC interface {
	InsertPC(context.Context, PC) error
	FetchPC(context.Context, FilterPC) (PC, error)
	DeletePC(context.Context, FilterPC) error
}
