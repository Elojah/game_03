package app

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/faction"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
)

const (
	factionBatchSize   = 10
	waypointsBatchSize = 10
)

var _ room.App = (*App)(nil)

type App struct {
	room.Store
	room.StorePublic
	room.StoreCell
	room.StoreUser
	room.StoreWorld
	room.StoreWorldCell
	room.StoreWorldSpawn
	room.StoreWorldWaypoint

	Entity  entity.App
	Faction faction.App
}

// PopulateWaypoints add spawns to all waypoints.
func (a App) PopulateWaypoints(ctx context.Context, ws room.WorldWaypoints) error {
	if a.Entity == nil {
		return errors.ErrMissingImplementation{Interface: "entity"}
	}

	// #Create and add altar in each waypoint
	altar, err := a.Entity.FetchTemplate(ctx, entity.FilterTemplate{
		// Pick first result (random) of altar template
		// TODO: Add as parameter to pick specific Altar
		Name: func(s string) *string { return &s }("Altar"),
	})
	if err != nil {
		return err
	}

	eAltar, err := a.Entity.Fetch(ctx, entity.Filter{
		ID: altar.EntityID,
	})
	if err != nil {
		return err
	}

	ans, _, err := a.Entity.FetchManyAnimation(ctx, entity.FilterAnimation{
		EntityID: altar.EntityID,
		Size:     1, // TODO: fetch all when more anims available
	})
	if err != nil {
		return err
	}

	if len(ans) == 0 {
		err := errors.ErrMissingDefaultAnimations{EntityID: altar.EntityID.String()}

		return err
	}

	for _, w := range ws {
		entityID := ulid.NewID()

		var main ulid.ID

		for _, an := range ans {
			an.ID = ulid.NewID()
			an.EntityID = entityID

			// default animation name is always "main"
			if an.Name == "main" {
				main = an.ID
			}

			if err := a.Entity.InsertAnimation(ctx, an); err != nil {
				return err
			}
		}

		// #Insert entity
		e := entity.E{
			ID:          entityID,
			UserID:      ulid.NewID(),
			CellID:      w.CellID,
			X:           w.X,
			Y:           w.Y,
			Rot:         0,
			Radius:      eAltar.Radius,
			At:          time.Now().UnixNano(),
			AnimationID: main,
			AnimationAt: 0,
			Objects:     eAltar.Objects,
		}
		if err := a.Entity.Insert(ctx, e); err != nil {
			return err
		}

		// #Associate entity as world spawn
		if err := a.StoreWorldSpawn.InsertWorldSpawn(ctx, room.WorldSpawn{
			ID:       ulid.NewID(),
			EntityID: e.ID,
			WorldID:  w.WorldID,
			OwnerID:  ulid.NewZeroID(),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (a App) CopyWorld(ctx context.Context, worldID ulid.ID) (ulid.ID, error) {
	copy, err := a.FetchWorld(ctx, room.FilterWorld{
		ID: worldID,
	})
	if err != nil {
		return nil, err
	}

	// Create world with new ID
	copy.ID = ulid.NewID()
	if err := a.InsertWorld(ctx, copy); err != nil {
		return nil, err
	}

	// Copy cells one by one
	var waypoints []room.WorldWaypoint

	for y := int64(0); y < copy.Height; y++ {
		for x := int64(0); x < copy.Width; x++ {
			wc, err := a.FetchWorldCell(ctx, room.FilterWorldCell{
				WorldID: worldID,
				X:       x,
				Y:       y,
			})
			if err != nil {
				return nil, err
			}

			c, err := a.FetchCell(ctx, room.FilterCell{
				ID: wc.CellID,
			})
			if err != nil {
				return nil, err
			}

			c.ID = ulid.NewID()
			if err := a.InsertCell(ctx, c); err != nil {
				return nil, err
			}

			if err := a.InsertWorldCell(ctx, room.WorldCell{
				WorldID: copy.ID,
				CellID:  c.ID,
				X:       wc.X,
				Y:       wc.Y,
			}); err != nil {
				return nil, err
			}

			// Fetch & Copy waypoints
			wps, _, err := a.FetchManyWorldWaypoint(ctx, room.FilterWorldWaypoint{
				CellID: wc.CellID,
				Size:   waypointsBatchSize,
			})
			if err != nil {
				return nil, err
			}

			for _, wp := range wps {
				wp.CellID = c.ID
				wp.WorldID = copy.ID

				if err := a.InsertWorldWaypoint(ctx, wp); err != nil {
					return nil, err
				}

				waypoints = append(waypoints, wp)
			}
		}
	}

	if err := a.PopulateWaypoints(ctx, waypoints); err != nil {
		return nil, err
	}

	// Fetch & Copy factions
	factions, _, err := a.Faction.FetchMany(ctx, faction.Filter{
		WorldID: worldID,
		Size:    factionBatchSize,
	})
	if err != nil {
		return nil, err
	}

	for _, fac := range factions {
		fac.WorldID = copy.ID

		if err := a.Faction.Insert(ctx, fac); err != nil {
			return nil, err
		}
	}

	return copy.ID, nil
}
