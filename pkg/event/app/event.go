package app

import (
	"context"
	"strconv"

	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/event"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
)

const (
	bufferSize = 100
)

type App struct {
	event.Cache
	event.CacheQ

	Entity  entity.App
	Ability ability.App
}

func (a App) Listen(ctx context.Context, entityID ulid.ID) error {
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	logger := log.With().Str("method", "subscribe").Str("entity", entityID.String()).Logger()

	// Subscribe is blocking
	err := a.CacheQ.Subscribe(ctx, event.FilterQ{
		EntityID: entityID,
	}, func(m event.Message) {
		var ev event.E

		if err := ev.Unmarshal([]byte(m.Message)); err != nil {
			log.Error().Err(err).Msg("failed to read message")

			return
		}

		if err := a.Cache.Insert(ctx, ev); err != nil {
			log.Error().Err(err).Msg("failed to insert event")

			return
		}

		at := strconv.FormatInt(ev.At, 10)

		// Fetch last entity state
		lasts, err := a.Entity.FetchManyCache(ctx, entity.FilterCache{
			ID:      entityID,
			Min:     "0",
			Max:     at,
			Size:    1,
			Reverse: true,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to fetch entity")

			return
		}

		if len(lasts) != 1 {
			err := errors.ErrNotFound{Resource: "entity", Index: entityID.String()}
			logger.Error().Err(err).Msg("no entity found")

			return
		}

		e := lasts[0]

		// Delete all future entities
		if err := a.Entity.DeleteCache(ctx, entity.FilterCache{
			ID:   entityID,
			Min:  at,
			Max:  "+inf",
			Size: -1,
		}); err != nil {
			logger.Error().Err(err).Msg("failed to delete entities")

			return
		}

		// Eval all events since at (should always include current one)
		events, err := a.Cache.FetchMany(ctx, event.Filter{
			EntityID: entityID,
			Min:      at,
			Max:      "+inf",
			Size:     -1,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to fetch events")

			return
		}

		logger.Info().Int("events", len(events)).Msg("eval events")

		for _, ev := range events {
			// if eval fail and no modifications, maybe don't process all above steps ?
			e, events, err := a.apply(ctx, ev, e)
			if err != nil {
				logger.Error().Err(err).Msg("failed to eval event to entity")

				continue
			}

			if err := a.Entity.Insert(ctx, e); err != nil {
				logger.Error().Err(err).Msg("failed to insert cache entity")

				continue
			}

			// trigger chain events
			for _, e := range events {
				if err := a.Publish(ctx, e); err != nil {
					logger.Error().Err(err).Msg("failed to publish event")

					continue
				}
			}
		}
	})

	return err
}

// TODO: everything in this function, check validity of received cast
func (a App) apply(ctx context.Context, ev event.E, e entity.E) (entity.E, map[string]event.E, error) {
	// self cast
	if ev.Effect.Self && ev.Effect.AbilityID.IsZero() {
		return a.move(ctx, ev, e)
	} else if ev.Effect.Self && !ev.Effect.AbilityID.IsZero() {
		return a.cast(ctx, ev, e)
	}

	targetType, ok := ev.Effect.Effect.Targets[ev.Effect.CurrentID.String()]
	if !ok {
		// no go case, current target is not specified in ability targets
		// so we don't know how to check target validity
		return e, nil, nil
	}

	// #Check cast/target validity
	var events map[string]event.E

	target, ok := ev.Effect.Targets[ev.Effect.CurrentID.String()]
	if !ok {
		return e, events, nil
	}

	// check target type validity
	// TODO: geometry not considered yet
	if targetType.Type == ability.Foe && e.FactionID.Compare(ev.Source.FactionID) == 0 {
		return e, nil, errors.ErrInvalidCastTargetFaction{
			AbilityID:        ev.Effect.AbilityID.String(),
			EffectID:         ev.Effect.EffectID.String(),
			EffectTargetID:   ev.Effect.CurrentID.String(),
			EffectTargetType: targetType.Type.String(),
			SourceID:         ev.Source.ID.String(),
			SourceFactionID:  ev.Source.FactionID.String(),
			TargetID:         e.ID.String(),
			TargetFactionID:  e.FactionID.String(),
		}
	}

	if targetType.Type == ability.Ally && e.FactionID.Compare(ev.Source.FactionID) != 0 {
		return e, nil, errors.ErrInvalidCastTargetFaction{
			AbilityID:        ev.Effect.AbilityID.String(),
			EffectID:         ev.Effect.EffectID.String(),
			EffectTargetID:   ev.Effect.CurrentID.String(),
			EffectTargetType: targetType.Type.String(),
			SourceID:         ev.Source.ID.String(),
			SourceFactionID:  ev.Source.FactionID.String(),
			TargetID:         e.ID.String(),
			TargetFactionID:  e.FactionID.String(),
		}
	}

	if targetType.Type == ability.Foe || targetType.Type == ability.Ally {
		if target.ID.Compare(e.ID) != 0 {
			return e, nil, errors.ErrInvalidTarget{
				AbilityID:      ev.Effect.AbilityID.String(),
				EffectID:       ev.Effect.EffectID.String(),
				EffectTargetID: ev.Effect.CurrentID.String(),
				SourceID:       ev.Source.ID.String(),
				TargetID:       target.ID.String(),
				EntityID:       e.ID.String(),
			}
		}

		if targetType.Range < int64(e.Distance(ev.Source)) {
			return e, nil, errors.ErrOutOfRangeTarget{
				AbilityID:      ev.Effect.AbilityID.String(),
				EffectID:       ev.Effect.EffectID.String(),
				EffectTargetID: ev.Effect.CurrentID.String(),
				SourceID:       ev.Source.ID.String(),
				TargetID:       target.ID.String(),
				TargetX:        float64(e.X),
				TargetY:        float64(e.Y),
				SourceX:        float64(ev.Source.X),
				SourceY:        float64(ev.Source.Y),
			}
		}

		// Direct amount on stats
		// TODO: EVERYTHING
		e.SetStat(ev.Effect.Effect.Stat, e.GetStat(ev.Effect.Effect.Stat)+ev.Effect.Effect.Amount.Direct)
	} else {
		return e, nil, errors.ErrNotImplemented{Version: "1.0.0"}
	}

	return e, events, nil
}

func (a App) move(ctx context.Context, ev event.E, e entity.E) (entity.E, map[string]event.E, error) {
	targetType, ok := ev.Effect.Effect.Targets[ev.Effect.CurrentID.String()]
	if !ok {
		// no go case, current target is not specified in ability targets
		// so we don't know how to check target validity
		return e, nil, nil
	}

	// self walk
	if m := targetType.MoveTarget; ev.Effect.AbilityID.Compare(ev.Source.ID) == 0 &&
		m.Move == ability.Walk &&
		m.TargetType == ability.Circle {
		// standard move (use center)
		pos, ok := ev.Effect.Targets[targetType.MoveTarget.TargetID.String()]
		if !ok {
			// position not defined, ignore
			return e, nil, nil
		}

		// TODO: check distance depending on current speed

		e.X = pos.Circle.X
		e.Y = pos.Circle.Y
	}

	return e, nil, nil
}

// cast implements self cast animations and target events propagation.
func (a App) cast(ctx context.Context, ev event.E, e entity.E) (entity.E, map[string]event.E, error) {
	// self cast
	ab, err := a.Ability.Fetch(ctx, ability.Filter{
		ID: ev.Effect.AbilityID,
	})
	if err != nil {
		return e, nil, err
	}

	// Fetch entity ability
	ea, err := a.Entity.FetchAbility(ctx, entity.FilterAbility{
		EntityID:  e.ID,
		AbilityID: ab.ID,
	})
	if err != nil {
		return e, nil, err
	}

	// #Check cast validity
	// TODO: consider mana buffs/debuffs here ?

	// check mana cost
	if e.Stats.MP < ab.ManaCost {
		return e, nil, errors.ErrNotEnoughMP{
			Required:  ab.ManaCost,
			MP:        e.Stats.MP,
			AbilityID: ab.ID.String(),
		}
	}

	// check cooldown
	if ev.At-ea.LastCast <
		(ab.Cooldown * (1 - (e.Stats.CooldownReduction / 100))) { //nolint: gomnd
		return e, nil, errors.ErrCooldownInProgress{
			At:        ev.At,
			LastCast:  ea.LastCast,
			Cooldown:  ab.Cooldown,
			AbilityID: ab.ID.String(),
		}
	}

	// set ability animation for casting entity
	if e.Abilities == nil {
		e.Abilities = make(map[string]entity.AnimationAbility)
	}

	e.Abilities[ulid.NewID().String()] = entity.AnimationAbility{
		AnimationID: ab.Animation,
	}

	var events map[string]event.E

	// create target events
	for effectID, ef := range ab.Effects {
		effectID := ulid.MustParse(effectID)

		// TODO: check triggers/modifiers
		for key, target := range ef.Targets {
			targetID, err := ulid.Parse(key)
			if err != nil {
				return e, nil, err
			}

			t, ok := ev.Effect.Targets[key]
			if !ok {
				// target not found (always optional ?)
				continue
			}

			// TODO: Fill geometry options
			// Check target validity and assign corresponding events
			switch target.Type {
			case ability.NoneTarget:
				// ignore ?
			case ability.Self:
				// target is not self
				if t.ID.Compare(e.ID) != 0 {
					return e, nil, errors.ErrInvalidSelfTarget{
						AbilityID:      ab.ID.String(),
						EffectID:       effectID.String(),
						EffectTargetID: key,
						SourceID:       e.ID.String(),
						TargetID:       t.ID.String(),
					}
				}

				// Do we really want to create a new event here ?
				// Apply effect directly ?
				id := ulid.NewID()
				events[id.String()] = event.E{
					ID:       id,
					EntityID: t.ID,
					Source:   e,
					At:       ev.At,
					Effect: ability.CastEffect{
						AbilityID: ab.ID,
						EffectID:  effectID,
						Effect:    ef,
						Targets: map[string]ability.CastTarget{
							key: {ID: t.ID}, // confirmed same as sourceID (self) above
						},
						CurrentID: targetID,
					},
				}
			case ability.Foe:
				id := ulid.NewID()
				events[id.String()] = event.E{
					ID:       id,
					EntityID: t.ID,
					Source:   e,
					At:       ev.At,
					Effect: ability.CastEffect{
						AbilityID: ab.ID,
						EffectID:  effectID,
						Effect:    ef,
						Targets: map[string]ability.CastTarget{
							key: {ID: t.ID},
						},
						CurrentID: targetID,
					},
				}
			case ability.Ally:
				id := ulid.NewID()
				events[id.String()] = event.E{
					ID:       id,
					EntityID: t.ID,
					Source:   e,
					At:       ev.At,
					Effect: ability.CastEffect{
						AbilityID: ab.ID,
						EffectID:  effectID,
						Effect:    ef,
						Targets: map[string]ability.CastTarget{
							key: {ID: t.ID},
						},
						CurrentID: targetID,
					},
				}
			case ability.Rect:
				return e, nil, errors.ErrNotImplemented{Version: "1.0.0"}
			case ability.Circle:
				return e, nil, errors.ErrNotImplemented{Version: "1.0.0"}
			}
		}
	}

	return e, events, nil
}
