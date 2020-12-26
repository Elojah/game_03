package space

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterWorld struct {
	ID *ulid.ID
}

type StoreWorld interface {
	UpsertWorld(context.Context, World) error
	FetchWorld(context.Context, FilterWorld) (World, error)
	FetchManyWorld(context.Context, FilterWorld, func(World) error) error
	DeleteWorld(context.Context, FilterWorld) error
	DeleteManyWorld(context.Context, FilterWorld) error
}
