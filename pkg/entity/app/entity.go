package app

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/ulid"
)

var _ entity.App = (*App)(nil)

type App struct {
	entity.Cache
	entity.Store
	entity.CacheAbility
	entity.StoreAbility
	entity.StoreAnimation
	entity.StoreBackup
	entity.StorePC
	entity.StoreTemplate
	entity.StoreSpawn
}

func (a App) CreateEntityFromBackup(ctx context.Context, id ulid.ID) (entity.E, error) {
	e, err := a.FetchBackup(ctx, entity.FilterBackup{
		ID: id,
	})
	if err != nil {
		return entity.E{}, err
	}

	e.At = time.Now().UnixMilli()
	if err := a.Insert(ctx, e); err != nil {
		return entity.E{}, err
	}

	if err := a.InsertCache(ctx, e); err != nil {
		return entity.E{}, err
	}

	// #Insert cache ability
	abs, _, err := a.FetchManyAbility(ctx, entity.FilterAbility{
		EntityID: id,
		Size:     entity.MaxAbilities,
	})
	if err != nil {
		return entity.E{}, err
	}

	for _, ab := range abs {
		if err := a.InsertCacheAbility(ctx, ab); err != nil {
			return entity.E{}, err
		}
	}

	return e, nil
}

func (a App) CreateBackupFromEntity(ctx context.Context, id ulid.ID) error {
	// #Fetch entity
	e, err := a.Fetch(ctx, entity.Filter{
		ID: id,
	})
	if err != nil {
		return err
	}

	// #Insert backup entity
	e.At = time.Now().UnixMilli()
	if err := a.UpdateBackup(ctx, entity.Filter{
		ID: id,
	}, e); err != nil {
		return err
	}

	// #Delete entity cache
	if err := a.DeleteCache(ctx, entity.FilterCache{
		ID: id,
	}); err != nil {
		return err
	}

	// #Delete entity
	if err := a.Delete(ctx, entity.Filter{
		ID: id,
	}); err != nil {
		return err
	}

	// #Delete cache ability
	abs, _, err := a.FetchManyAbility(ctx, entity.FilterAbility{
		EntityID: id,
		Size:     entity.MaxAbilities,
	})
	if err != nil {
		return err
	}

	for _, ab := range abs {
		if err := a.DeleteCacheAbility(ctx, entity.FilterAbility{
			AbilityID: ab.AbilityID,
		}); err != nil {
			return err
		}
	}

	return nil
}
