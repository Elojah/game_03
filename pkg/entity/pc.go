package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterPC struct {
	ID     *ulid.ID
	UserID *ulid.ID
	RoomID *ulid.ID
}

type StorePC interface {
	InsertPC(context.Context, PC) error
	FetchPC(context.Context, FilterPC) (PC, error)
	FetchManyPC(context.Context, FilterPC) ([]PC, error)
	DeletePC(context.Context, FilterPC) error
}

type PCs []PC

func (pcs PCs) RoomIDs() []ulid.ID {
	result := make([]ulid.ID, 0, len(pcs))

	for _, pc := range pcs {
		result = append(result, pc.RoomID)
	}

	return result
}
