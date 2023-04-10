package ability

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/ulid"
)

var (
	tmpFetchEntity  = func(ulid.ID) (entity.E, error) { return entity.E{}, nil }
	tmpUpdateEntity = func(entity.E) error { return nil }

	operators = map[Operator]func(a int64, b int64) bool{
		NoneOperator: func(a, b int64) bool { return false },
		Equal:        func(a, b int64) bool { return a == b },
		NotEqual:     func(a, b int64) bool { return a != b },
		Greater:      func(a, b int64) bool { return a > b },
		Lesser:       func(a, b int64) bool { return a < b },
	}
)

func (a Ability) Cast(cast Cast) error {
	for compID, comp := range a.Components {
		for _, trigger := range comp.Triggers {
			abilityM, effectM := trigger.Eval()
		}

		for effectID, effect := range comp.Effects {
			cts, ok := cast.Targets[effectID]
			if !ok {
				continue
			}

			for _, t := range cts.Targets {
				/*
					check validity of targets (number, type, range)
				*/
				e, err := tmpFetchEntity(t.ID)
				if err != nil {
					// compose error ?
				}

				ms, err := effect.Apply(e, cast)
				if err != nil {
					// compose error ?
				}

				if err := a.ApplyModifiers(compID, ms); err != nil {
					// compose error ?
				}

			}
		}
	}

	return nil
}

func (a *Ability) ApplyModifiers(compID string, ms []Modifier) error {
	for _, m := range ms {
		a.CastTime += m.CastTime
		a.ManaCost += m.ManaCost
		a.Cooldown += m.Cooldown

		comp, ok := a.Components[compID]
		if !ok {
			// return err
		}

		effectID := m.Amount.EffectID.String()
		if effect, ok := comp.Effects[effectID]; ok {
			if m.Cancel {
				// TODO ?
			}

			// modifier rules, change only in a first step
			effect.Amount = m.Amount
			effect.Drain = m.Drain
			effect.Duration = m.Duration
			effect.Delay = m.Delay
			effect.Repeat = m.Repeat
			effect.StackRules = m.StackRules

			a.Components[compID].Effects[effectID] = effect
		}

	}

	return nil
}

func (effect Effect) Apply(e entity.E, c Cast) ([]Modifier, error) {
	var result []Modifier

	for _, trigger := range effect.Triggers {
		trigger.Eval()
		// apply on effects ?
	}

	return result, nil
}

func (t Trigger) Eval(c Cast) (bool, error) {
	amount, err := t.Amount.Compute(c)
	if err != nil {
		return false, err
	}

	treshold, err := t.Treshold.Compute(c)
	if err != nil {
		return false, err
	}

	comp, ok := operators[t.Operator]
	if !ok {
		return false, errors.ErrInvalidOperator{Value: t.Operator.String()}
	}

	if !comp(amount, treshold) {
		// trigger not fulfilled
		return false, nil
	}

	return true, nil
}

func (amount Amount) Compute(c Cast) (int64, error) {
	result := amount.Direct

	if amount.Percentage == 0 && amount.EffectID.IsZero() {
		return result, nil
	}

	targets, ok := c.Targets[amount.ID.String()]
	if !ok || len(targets.Targets) == 0 {
		return result, nil // error ?
	}

	e, err := tmpFetchEntity(targets.Targets[0].ID)
	if err != nil {
		return result, err
	}

	if !amount.EffectID.IsZero() {
		n, ok := e.Effects[amount.EffectID.String()]
		if !ok {
			return result, nil // no error, effect is not present on target
		}

		return result + n, nil
	}

	pc := e.FetchStat(amount.Stat) * amount.Percentage

	return amount.Direct + pc, nil
}

type Store interface{}

type App interface{}
