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

// CreateFromCast creates events from cast.
func (a App) CreateFromCast(ctx context.Context, sourceID ulid.ID, c ability.Cast) (map[string]event.E, error) {
	logger := log.With().Str("method", "create_from_cast").Str("source", sourceID.String()).Int64("at", c.At).Logger()

	result := make(map[string]event.E, 1)

	// Fetch last entity state
	lasts, err := a.Entity.FetchManyCache(ctx, entity.FilterCache{
		ID:      sourceID,
		Min:     "0",
		Max:     strconv.FormatInt(c.At, 10),
		Size:    1,
		Reverse: true,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch entity")

		return nil, err
	}

	if len(lasts) == 0 {
		err := errors.ErrNotFound{Resource: "entity", Index: sourceID.String()}

		logger.Error().Err(err).Msg("failed to fetch entity")

		return nil, err
	}

	source := lasts[0]

	// #Sepcial case: abilityID = 0 -> use this case for move events
	if target, ok := c.Targets[sourceID.String()]; ok && c.AbilityID.IsZero() {
		position := ability.CastTarget{
			ID:     ulid.NewID(),
			Circle: target.Circle,
		}
		ev := event.E{
			ID:       ulid.NewID(),
			EntityID: sourceID,
			Source:   source,
			At:       c.At,
			Effect: ability.CastEffect{
				// default move effect, hard code here
				Effect: ability.Effect{
					Targets: map[string]ability.Target{
						sourceID.String(): {
							Type: ability.Self,
							Move: ability.MoveTarget{
								Move:       ability.Walk,
								TargetType: ability.Circle,
								TargetID:   position.ID,
							},
						},
					},
				},
				Targets: map[string]ability.CastTarget{
					position.ID.String(): position,
				},
				CurrentID: sourceID,
			},
		}

		result[ev.ID.String()] = ev

		return result, nil
	}

	// Fetch ability
	ab, err := a.Ability.Fetch(ctx, ability.Filter{ID: c.AbilityID})
	if err != nil {
		return nil, err
	}

	// Fetch entity ability
	ea, err := a.Entity.FetchAbility(ctx, entity.FilterAbility{
		EntityID:  source.ID,
		AbilityID: ab.ID,
	})
	if err != nil {
		return nil, err
	}

	// #Check cast validity
	// TODO: consider mana buffs/debuffs here ?

	// check mana cost
	if source.Stats.MP < ab.ManaCost {
		return nil, errors.ErrNotEnoughMP{
			Required:  ab.ManaCost,
			MP:        source.Stats.MP,
			AbilityID: ab.ID.String(),
		}
	}

	// check cooldown
	if c.At-ea.LastCast <
		(ab.Cooldown * (1 - (source.Stats.CooldownReduction / 100))) {
		return nil, errors.ErrCooldownInProgress{
			At:        c.At,
			LastCast:  ea.LastCast,
			Cooldown:  ab.Cooldown,
			AbilityID: ab.ID.String(),
		}
	}

	// Create source cast event for source
	ev := event.E{
		ID:         ulid.NewID(),
		EntityID:   sourceID,
		Source:     source,
		At:         c.At,
		SourceCast: c,
	}

	result[ev.ID.String()] = ev

	for effectID, ef := range ab.Effects {
		// TODO: check triggers/modifiers
		for key, target := range ef.Targets {
			targetID, err := ulid.Parse(key)
			if err != nil {
				return nil, err
			}

			t, ok := c.Targets[key]
			if !ok {
				continue
			}

			// TODO: Fill geometry options
			// Check target validity and assign corresponding events
			switch target.Type {
			case ability.NoneTarget:
				// ignore ?
			case ability.Self:
				// target is not self
				if t.ID.Compare(sourceID) != 0 {
					return nil, errors.ErrInvalidSelfTarget{
						AbilityID:      ab.ID.String(),
						EffectID:       effectID,
						EffectTargetID: key,
						SourceID:       sourceID.String(),
						TargetID:       t.ID.String(),
					}
				}

				id := ulid.NewID()
				result[id.String()] = event.E{
					ID:       id,
					EntityID: t.ID,
					Source:   source,
					At:       c.At,
					Effect: ability.CastEffect{
						Effect: ef,
						Targets: map[string]ability.CastTarget{
							key: {ID: t.ID}, // confirmed same as sourceID (self) above
						},
						CurrentID: targetID,
					},
				}
			case ability.ClosestSelf:
			case ability.Foe:
				id := ulid.NewID()
				result[id.String()] = event.E{
					ID:       id,
					EntityID: t.ID,
					Source:   source,
					At:       c.At,
					Effect: ability.CastEffect{
						Effect: ef,
						Targets: map[string]ability.CastTarget{
							key: {ID: t.ID},
						},
						CurrentID: targetID,
					},
				}
			case ability.ClosestFoe:
			case ability.Ally:
				id := ulid.NewID()
				result[id.String()] = event.E{
					ID:       id,
					EntityID: t.ID,
					Source:   source,
					At:       c.At,
					Effect: ability.CastEffect{
						Effect: ef,
						Targets: map[string]ability.CastTarget{
							key: {ID: t.ID},
						},
						CurrentID: targetID,
					},
				}
			case ability.ClosestAlly:
			case ability.Rect:
			case ability.Circle:
			}
		}
	}

	return result, nil
}

func (a App) Eval(ctx context.Context, entityID ulid.ID) error {
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
			// add cast state to entity if event is cast
			e, err = a.Ability.AddCast(ctx, e, ev.SourceCast)
			if err != nil {
				logger.Error().Err(err).Msg("failed to add ability cast to entity")

				continue
			}

			e = a.EvalEntity(ctx, ev, e)

			if err := a.Entity.Insert(ctx, e); err != nil {
				logger.Error().Err(err).Msg("failed to insert cache entity")

				continue
			}

			// trigger chain events ?
		}
	})

	return err
}

// TODO: everything in this function, check validity of received cast
func (a App) EvalEntity(ctx context.Context, ev event.E, e entity.E) entity.E {
	current := ev.Effect.CurrentID.String()

	tt, ok := ev.Effect.Effect.Targets[current]
	if !ok {
		// no go case, current target is not specified in ability targets
		// so we don't know which ability's target to pick
		return e
	}

	// self move first
	if m := tt.Move; m.Move != ability.NoneMove {
		switch m.TargetType {
		case ability.NoneTarget:
			// no move destination ? ignore
		case ability.Self:
			// move on self ? ignore
		case ability.ClosestSelf:
			// chain event
		case ability.Foe:
			// chain event
		case ability.ClosestFoe:
			// chain event
		case ability.Ally:
			// chain event
		case ability.ClosestAlly:
			// chain event
		case ability.Rect:
			// move to rect ? ignore, use circle instead
		case ability.Circle:
			// standard move (use center)
			pos, ok := ev.Effect.Targets[m.TargetID.String()]
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
