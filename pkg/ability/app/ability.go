package app

import (
	"context"

	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/ulid"
)

type App struct {
	ability.Cache
	ability.Store
}

func (a App) AddCast(ctx context.Context, e entity.E, cast ability.Cast) (entity.E, error) {
	if cast.AbilityID.IsZero() {
		return e, nil
	}

	if e.Abilities == nil {
		e.Abilities = make(map[string]entity.AnimationAbility)
	}

	ab, err := a.Fetch(ctx, ability.Filter{
		ID: cast.AbilityID,
	})
	if err != nil {
		return e, err
	}

	e.Abilities[ulid.NewID().String()] = entity.AnimationAbility{
		AnimationID: ab.Animation,
	}

	return e, nil
}
