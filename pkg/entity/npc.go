package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterNPC struct {
	ID      ulid.ID
	IDs     []ulid.ID
	WorldID ulid.ID

	State []byte
	Size  int
}

type StoreNPC interface {
	InsertNPC(context.Context, NPC) error
	FetchNPC(context.Context, FilterNPC) (NPC, error)
	FetchManyNPC(context.Context, FilterNPC) ([]NPC, []byte, error)
	DeleteNPC(context.Context, FilterNPC) error
}
