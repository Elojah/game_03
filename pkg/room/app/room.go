package app

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
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

	Entity entity.App
}

// PopulateWaypoints
func (a App) PopulateWaypoints(ctx context.Context, ws room.Waypoints, worldID ulid.ID) error {
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
			X:           w.Position.X,
			Y:           w.Position.Y,
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
			WorldID:  worldID,
			OwnerID:  ulid.NewZeroID(),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (a App) CopyWorld(ctx context.Context, worldID ulid.ID) (ulid.ID, error) {
	orig, err := a.FetchWorld(ctx, room.FilterWorld{
		ID: worldID,
	})
	if err != nil {
		return nil, err
	}

	// Create world with new ID
	orig.ID = ulid.NewID()
	if err := a.InsertWorld(ctx, orig); err != nil {
		return nil, err
	}

	// Copy cells one by one
	for y := int64(0); y < orig.Height; y++ {
		for x := int64(0); x < orig.Width; x++ {
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

			wc.WorldID = orig.ID
			wc.CellID = c.ID

			if err := a.InsertWorldCell(ctx, wc); err != nil {
				return nil, err
			}
		}
	}

	return orig.ID, nil
}
