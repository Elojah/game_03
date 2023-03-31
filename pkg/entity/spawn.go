package entity

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterSpawn struct {
	EntityID  ulid.ID
	EntityIDs []ulid.ID
	SpawnID   ulid.ID
	SpawnIDs  []ulid.ID

	All bool

	State []byte
	Size  int
}

type StoreSpawn interface {
	InsertSpawn(context.Context, Spawn) error
	FetchSpawn(context.Context, FilterSpawn) (Spawn, error)
	FetchManySpawn(context.Context, FilterSpawn) ([]Spawn, []byte, error)
	DeleteSpawn(context.Context, FilterSpawn) error
}
