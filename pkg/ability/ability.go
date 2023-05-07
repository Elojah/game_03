package ability

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/ulid"
)

var (
	tmpFetchEntity         = func(ulid.ID) (entity.E, error) { return entity.E{}, nil }
	tmpUpdateEntity        = func(entity.E) error { return nil }
	tmpFetchEntityAbility  = func(ulid.ID, ulid.ID) (entity.Ability, error) { return entity.Ability{}, nil }
	tmpUpdateEntityAbility = func(entity.Ability) error { return nil }

	operators = map[Operator]func(a int64, b int64) bool{
		NoneOperator: func(a, b int64) bool { return false },
		Equal:        func(a, b int64) bool { return a == b },
		NotEqual:     func(a, b int64) bool { return a != b },
		Greater:      func(a, b int64) bool { return a > b },
		Lesser:       func(a, b int64) bool { return a < b },
	}
)

// func (a Ability) Cast(cast Cast) error {
// 	now := time.Now()

// 	// Check if entity is casting
// 	e, err := tmpFetchEntity(cast.SourceID) // now
// 	if err != nil {
// 		return err
// 	}

// 	// check if entity owns ability + last cast
// 	ea, err := tmpFetchEntityAbility(cast.SourceID, a.ID)
// 	if err != nil {
// 		return err
// 	}

// 	// TODO: Cast time how to manage ?
// 	// Cast effects
// 	a.Eval(cast)

// 	// Check mana cost and cooldown
// 	// TODO: consider mana buffs/debuffs here ?
// 	if e.Stats.MP < a.ManaCost {
// 		return errors.ErrNotEnoughMP{
// 			Required:  a.ManaCost,
// 			MP:        e.Stats.MP,
// 			AbilityID: a.ID.String(),
// 		}
// 	}

// 	// TODO: consider cooldown buffs/debuffs here ?
// 	// # Formula Cooldown
// 	if now.UnixMilli()-ea.LastCast <
// 		(a.Cooldown * (1 - (e.Stats.CooldownReduction / 100))) {
// 		return errors.ErrCooldownInProgress{
// 			At:        now.UnixMilli(),
// 			LastCast:  ea.LastCast,
// 			Cooldown:  a.Cooldown,
// 			AbilityID: a.ID.String(),
// 		}
// 	}

// 	return nil
// }

// func (a *Ability) Eval(cast Cast) ([]entity.E, error) {
// 	// entities := make(map[string]entity.E)

// 	// for _, effect := range a.Effects {
// 	// 	es, am, err := effect.Eval(cast)
// 	// }

// 	// if err := tmpUpdateEntity(e); err != nil {
// 	// }

// 	return nil, nil
// }

// func (ef *Effect) Eval(c Cast) (map[string]entity.E, map[string]AbilityModifier, error) {
// 	am := make(map[string]AbilityModifier)

// 	// Evaluate triggers and modify
// 	for _, trigger := range ef.Triggers {
// 		ok, err := trigger.Eval(c)
// 		if err != nil {
// 			return nil, nil, err
// 		}

// 		if !ok {
// 			continue
// 		}

// 		for key, modifier := range trigger.AbilityModifiers {
// 			am[key] = modifier
// 		}

// 		for key, m := range trigger.EffectModifiers {
// 			// apply on self
// 			if key == "" {
// 				ef.Modify(m)
// 			} else if child, ok := ef.Effects[key]; ok {
// 				child.Modify(m)

// 				// reassign in map
// 				ef.Effects[key] = child
// 			}
// 		}
// 	}

// 	for tid, t := range ef.Targets {
// 		ct, ok := c.Targets[tid]
// 		if !ok {
// 			// no defined target, pass ?
// 		}

// 		e, err := tmpFetchEntity(ct.ID)
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 	}

// 	return nil, am, nil
// }

// func (ef *Effect) Modify(m EffectModifier) {
// 	ef.Stat = m.Stat
// 	ef.Amount = m.Amount
// 	ef.Duration += m.Duration
// 	ef.Delay += m.Delay
// 	ef.Repeat += m.Repeat
// 	ef.StackRules = m.StackRules
// }

// func (a *Ability) Modify(ms map[string]AbilityModifier) {
// 	for _, m := range ms {
// 		a.CastTime += m.CastTime
// 		a.ManaCost += m.ManaCost
// 		a.Cooldown += m.Cooldown
// 	}
// }

// func (t Trigger) Eval(c Cast) (bool, error) {
// 	amount, err := t.Amount.Eval(c)
// 	if err != nil {
// 		return false, err
// 	}

// 	treshold, err := t.Treshold.Eval(c)
// 	if err != nil {
// 		return false, err
// 	}

// 	comp, ok := operators[t.Operator]
// 	if !ok {
// 		return false, errors.ErrInvalidOperator{Value: t.Operator.String()}
// 	}

// 	if !comp(amount, treshold) {
// 		// trigger not raised
// 		return false, nil
// 	}

// 	return true, nil
// }

// func (amount Amount) Eval(c Cast) (int64, error) {
// 	result := amount.Direct

// 	if amount.Percentage == 0 && amount.EffectID.IsZero() {
// 		return result, nil
// 	}

// 	target, ok := c.Targets[amount.ID.String()]
// 	if !ok {
// 		return result, nil // error ?
// 	}

// 	e, err := tmpFetchEntity(target.ID)
// 	if err != nil {
// 		return result, err
// 	}

// 	if !amount.EffectID.IsZero() {
// 		n, ok := e.Effects[amount.EffectID.String()]
// 		if !ok {
// 			return result, nil // no error, effect is not present on target
// 		}

// 		return result + n, nil
// 	}

// 	pc := e.FetchStat(amount.Stat) * amount.Percentage

// 	return amount.Direct + pc, nil
// }

type Filter struct {
	ID  ulid.ID
	IDs []ulid.ID

	State []byte
	Size  int
}

type Store interface {
	Insert(context.Context, A) error
	Fetch(context.Context, Filter) (A, error)
	FetchMany(context.Context, Filter) ([]A, []byte, error)
	Delete(context.Context, Filter) error
}

// TODO
type Cache interface{}

type App interface {
	Cache
	Store
}

func (cef CastEffect) Eval(e entity.E) entity.E {
	current := cef.CurrentID.String()

	tt, ok := cef.Effect.Targets[current]
	if !ok {
		// no go case, current target is not specified in ability targets
		// so we don't know which ability's target to pick
		return e
	}

	// move first
	if tt.Move != NoneMove {
		switch tt.PositionTargetType {
		case NoneTarget:
			// no move destination ? ignore
		case Self:
			// move on self ? ignore
		case Foe:
			// chain event
		case ClosestSelf:
			// chain event
		case ClosestFoe:
			// chain event
		case Rect:
			// move to rect ? ignore
		case Circle:
			// standard move (use center)
			pos, ok := cef.Targets[tt.PositionTargetID.String()]
			if !ok {
				// position not defined, ignore
				break
			}

			e.X = pos.Circle.X
			e.Y = pos.Circle.Y
		}
	}

	return e
}
