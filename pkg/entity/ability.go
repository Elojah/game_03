package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterAbility struct {
	EntityID  ulid.ID
	AbilityID ulid.ID

	State []byte
	Size  int
}

type StoreAbility interface {
	InsertAbility(context.Context, Ability) error
	FetchAbility(context.Context, FilterAbility) (Ability, error)
	FetchManyAbility(context.Context, FilterAbility) ([]Ability, []byte, error)
	DeleteAbility(context.Context, FilterAbility) error
}

type CacheAbility interface {
	InsertCacheAbility(context.Context, Ability) error
	FetchCacheAbility(context.Context, FilterAbility) (Ability, error)
	DeleteCacheAbility(context.Context, FilterAbility) error
}
