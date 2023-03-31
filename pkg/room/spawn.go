package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterWorldSpawn struct {
	ID       ulid.ID
	IDs      []ulid.ID
	WorldID  ulid.ID
	WorldIDs []ulid.ID

	All bool

	State []byte
	Size  int
}

type StoreWorldSpawn interface {
	InsertWorldSpawn(context.Context, WorldSpawn) error
	FetchWorldSpawn(context.Context, FilterWorldSpawn) (WorldSpawn, error)
	FetchManyWorldSpawn(context.Context, FilterWorldSpawn) ([]WorldSpawn, []byte, error)
	DeleteWorldSpawn(context.Context, FilterWorldSpawn) error
}
