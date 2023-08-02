package app

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/ability"
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

	Ability ability.App
}

func (a App) Insert(ctx context.Context, e entity.E) error {
	if err := a.InsertCache(ctx, e); err != nil {
		return err
	}

	if err := a.Store.Insert(ctx, e); err != nil {
		return err
	}

	return nil
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

// CreateDefaultAbilities.
func (a App) CreateDefaultAbilities(ctx context.Context, entityID ulid.ID) error {
	// default ability: tp to spawn point
	// TODO: add animation
	abRes := ability.A{
		ID:        ulid.NewID(),
		Name:      "test_spell_res",
		Icon:      ulid.MustParse("01H0N2PHEMG0DJSDVGV15CXA3M"), // static wip icon
		Animation: nil,
		CastTime:  2000,
		ManaCost:  0,
		Cooldown:  10000,
		Effects: map[string]ability.Effect{
			ulid.NewID().String(): {
				Targets: map[string]ability.Target{
					ulid.NewID().String(): {
						GroupID: ulid.NewID(),
						Type:    ability.Self,
						MoveTarget: ability.MoveTarget{
							Move:       ability.Teleport,
							TargetType: ability.Spawn,
							TargetID:   ulid.NewID(),
						},
					},
				},
			},
		},
	}

	if err := a.Ability.Insert(ctx, abRes); err != nil {
		return err
	}

	if err := a.InsertAbility(ctx, entity.Ability{
		EntityID:  entityID,
		AbilityID: abRes.ID,
		LastCast:  0,
	}); err != nil {
		return err
	}

	// basic ability: damage one foe
	abilityTemplate, err := a.FetchTemplate(ctx, entity.FilterTemplate{
		Name: func(s string) *string { return &s }("ability"), // TODO: replace with "green_ability" template ?
	})
	if err != nil {
		return err
	}

	anim, err := a.FetchAnimation(ctx, entity.FilterAnimation{
		EntityID: abilityTemplate.EntityID,
	})
	if err != nil {
		return err
	}

	anim.ID = ulid.NewID()
	anim.EntityID = entityID

	ab := ability.A{
		ID:        ulid.NewID(),
		Name:      "test_spell",
		Icon:      ulid.MustParse("01H0N2PHEMG0DJSDVGV15CXA3M"), // static wip icon
		Animation: anim.ID,
		CastTime:  1000,
		ManaCost:  10,
		Cooldown:  3000,
		Effects: map[string]ability.Effect{
			ulid.NewID().String(): {
				Stat:   entity.HP,
				Amount: ability.Amount{Direct: -10},
				Targets: map[string]ability.Target{
					ulid.NewID().String(): {
						GroupID: ulid.NewID(),
						Type:    ability.Circle,
						Radius:  1,
						Range:   100,
					},
				},
			},
		},
	}

	if err := a.Ability.Insert(ctx, ab); err != nil {
		return err
	}

	if err := a.InsertAbility(ctx, entity.Ability{
		EntityID:  entityID,
		AbilityID: ab.ID,
		LastCast:  0,
	}); err != nil {
		return err
	}

	if err := a.InsertAnimation(ctx, anim); err != nil {
		return err
	}

	return nil
}
