package ability

import (
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/errors"
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

func (a Ability) Cast(cast Cast) error {
	now := time.Now()

	// Check if entity is casting
	e, err := tmpFetchEntity(cast.SourceID) // now
	if err != nil {
		return err
	}

	// Evaluate triggers and modify
	for _, trigger := range a.Triggers {
		ok, err := trigger.Eval(cast)
		if err != nil {
			// compose error ?
			return err
		}

		if ok {
			a.Modify(trigger.AbilityModifiers)
			a.ModifyEffect(trigger.EffectModifiers)
		}
	}

	// Check mana cost and cooldown
	// TODO: consider mana buffs/debuffs here ?
	if e.Stats.MP < a.ManaCost {
		return errors.ErrNotEnoughMP{
			Required:  a.ManaCost,
			MP:        e.Stats.MP,
			AbilityID: a.ID.String(),
		}
	}

	ea, err := tmpFetchEntityAbility(e.ID, a.ID)
	if err != nil {
		return err
	}

	// TODO: consider cooldown buffs/debuffs here ?
	// # Formula Cooldown
	if now.UnixMilli()-ea.LastCast <
		(a.Cooldown * (1 - (e.Stats.CooldownReduction / 100))) {
		return errors.ErrCooldownInProgress{
			At:        now.UnixMilli(),
			LastCast:  ea.LastCast,
			Cooldown:  a.Cooldown,
			AbilityID: a.ID.String(),
		}
	}

	// TODO: Cast time how to manage ?

	// Cast effects
	for effectID, effect := range a.Effects {
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
				return err
			}

		}
	}

	return nil
}

func (a *Ability) Modify(ms map[string]AbilityModifier) {
	for _, m := range ms {
		a.CastTime += m.CastTime
		a.ManaCost += m.ManaCost
		a.Cooldown += m.Cooldown
	}
}

func (a *Ability) ModifyEffect(ms map[string]EffectModifier) {
	for _, m := range ms {
		id := m.EffectID.String()

		e, ok := a.Effects[id]
		if !ok {
			continue
		}

		if m.Cancel {
			delete(a.Effects, id)
			continue
		}

		e.Stat = m.Stat
		e.Amount = m.Amount
		e.Drain = m.Drain

		e.Duration += m.Duration
		e.Delay += m.Delay
		e.Repeat += m.Repeat

		e.StackRules = m.StackRules

		a.Effects[id] = e
	}
}

func (t Trigger) Eval(c Cast) (bool, error) {
	amount, err := t.Amount.Eval(c)
	if err != nil {
		return false, err
	}

	treshold, err := t.Treshold.Eval(c)
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

func (amount Amount) Eval(c Cast) (int64, error) {
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
