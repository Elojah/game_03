package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type WorldWaypoints []WorldWaypoint

type FilterWorldWaypoint struct {
	ID     ulid.ID
	CellID ulid.ID

	State []byte
	Size  int
}

type StoreWorldWaypoint interface {
	InsertWorldWaypoint(context.Context, WorldWaypoint) error
	FetchWorldWaypoint(context.Context, FilterWorldWaypoint) (WorldWaypoint, error)
	DeleteWorldWaypoint(context.Context, FilterWorldWaypoint) error
	FetchManyWorldWaypoint(context.Context, FilterWorldWaypoint) ([]WorldWaypoint, []byte, error)
}
