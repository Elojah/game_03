package space

import (
	"context"

	geometry "github.com/elojah/game_03/pkg/geometry"
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

type FilterEntity struct {
	WorldID ulid.ID

	Min geometry.Vec2
	Max geometry.Vec2
	TS  uint64

	ID *ulid.ID
}

type CacheEntity interface {
	UpsertEntity(context.Context, Entity) error
	FetchEntity(context.Context, FilterEntity) (Entity, error)
	FetchManyEntity(context.Context, FilterEntity, func(Entity) error) error
	DeleteEntity(context.Context, FilterEntity) error
	DeleteManyEntity(context.Context, FilterEntity) error
}
