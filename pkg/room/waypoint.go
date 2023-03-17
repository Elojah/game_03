package room

import (
	"github.com/elojah/game_03/pkg/geometry"
	"github.com/elojah/game_03/pkg/ulid"
)

type Waypoint struct {
	Position geometry.Vec2
	CellID   ulid.ID
}

type Waypoints []Waypoint
