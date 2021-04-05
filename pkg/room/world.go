package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterWorld struct {
	ID *ulid.ID
}

type StoreWorld interface {
	InsertWorld(context.Context, World) error
	FetchWorld(context.Context, FilterWorld) (World, error)
	DeleteWorld(context.Context, FilterWorld) error
}