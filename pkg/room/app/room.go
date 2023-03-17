package app

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ room.App = (*App)(nil)

type App struct {
	room.Store
	room.StorePublic
	room.StoreCell
	room.StoreUser
	room.StoreWorld
	room.StoreWorldCell

	Entity entity.App
}

// PopulateWaypoints
func (a App) PopulateWaypoints(ctx context.Context, ws room.Waypoints) error {
	if a.Entity == nil {
		return errors.ErrMissingImplementation{Interface: "entity"}
	}

	altar, err := a.Entity.FetchTemplate(ctx, entity.FilterTemplate{
		// Pick first result (random) of altar template
		// TODO: Add as parameter to pick specific Altar
		Name: func(s string) *string { return &s }("Altar"),
	})
	if err != nil {
		return status.New(codes.Internal, err.Error()).Err()
	}

	ans, _, err := a.Entity.FetchManyAnimation(ctx, entity.FilterAnimation{
		EntityID: altar.ID,
		Size:     1, // TODO: fetch all when more anims available
	})
	if err != nil {
		return status.New(codes.Internal, err.Error()).Err()
	}

	if len(ans) == 0 {
		err := errors.ErrMissingDefaultAnimations{EntityID: altar.ID.String()}

		return status.New(codes.Internal, err.Error()).Err()
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
				return status.New(codes.Internal, err.Error()).Err()
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
			Radius:      10, //nolint: gomnd
			At:          time.Now().UnixNano(),
			AnimationID: main,
			AnimationAt: 0,
		}
		if err := a.Entity.Insert(ctx, e); err != nil {
			return status.New(codes.Internal, err.Error()).Err()
		}
	}

	return nil
}
