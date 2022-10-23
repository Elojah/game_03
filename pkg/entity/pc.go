package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterPC struct {
	ID      ulid.ID
	IDs     []ulid.ID
	UserID  ulid.ID
	WorldID ulid.ID

	State []byte
	Size  int
}

type StorePC interface {
	InsertPC(context.Context, PC) error
	FetchPC(context.Context, FilterPC) (PC, error)
	FetchManyPC(context.Context, FilterPC) ([]PC, []byte, error)
	DeletePC(context.Context, FilterPC) error
}
