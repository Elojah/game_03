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

func (a App) Eval(ctx context.Context, entityID ulid.ID) error {
	ch := make(chan int64, bufferSize)

	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	go func(ctx context.Context, entityID ulid.ID) {
		logger := log.With().Str("method", "subscribe").Str("entity", entityID.String()).Logger()

		for {
			select {
			case _ = <-ctx.Done():
				logger.Error().Err(ctx.Err()).Msg("context cancelled")

				return
			case m := <-ch:
				at := strconv.FormatInt(m, 10)

				// Fetch last entity state
				lasts, err := a.Entity.FetchManyCache(ctx, entity.FilterCache{
					ID:   entityID,
					Min:  "0",
					Max:  at,
					Size: 1,
				})
				if err != nil {
					logger.Error().Err(err).Msg("failed to fetch entity")

					continue
				}

				if len(lasts) != 1 {
					err := errors.ErrNotFound{Resource: "entity", Index: entityID.String()}
					logger.Error().Err(err).Msg("no entity found")

					continue
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

					continue
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

					continue
				}

				logger.Info().Int("events", len(events)).Msg("eval events")

				for _, ev := range events {
					e = ev.Eval(e)

					if err := a.Entity.InsertCache(ctx, e); err != nil {
						logger.Error().Err(err).Msg("failed to insert cache entity")

						continue
					}

					if err := a.Entity.Insert(ctx, e); err != nil {
						logger.Error().Err(err).Msg("failed to insert entity")

						continue
					}

					// trigger chain events ?
				}
			}
		}
	}(ctx, entityID)

	// Subscribe is blocking
	return a.CacheQ.Subscribe(ctx, event.FilterQ{
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

		ch <- ev.At
	})
}

// CreateFromCast creates events from cast.
func (a App) CreateFromCast(ctx context.Context, sourceID ulid.ID, c ability.Cast) (map[string]event.E, error) {
	result := make(map[string]event.E, 1)

	// we use this case for move events
	if target, ok := c.Targets[sourceID.String()]; ok && c.AbilityID.IsZero() {
		position := ability.CastTarget{
			ID:     ulid.NewID(),
			Circle: target.Circle,
		}
		ev := event.E{
			ID:       ulid.NewID(),
			EntityID: sourceID,
			SourceID: sourceID,
			At:       c.At,
			Effect: ability.CastEffect{
				// default move effect, hard code here
				Effect: ability.Effect{
					Targets: map[string]ability.Target{
						sourceID.String(): {
							Type:               ability.Self,
							Move:               ability.Walk,
							PositionTargetType: ability.Circle,
							PositionTargetID:   position.ID,
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

	ab, err := a.Ability.Fetch(ctx, ability.Filter{ID: c.AbilityID})
	if err != nil {
		return nil, err
	}

	for _, ef := range ab.Effects {
		// TODO: check triggers/modifiers
		for key, target := range ef.Targets {
			t, ok := c.Targets[key]
			if !ok {
				continue
			}

			// check target is same type than t
			_ = target

			// Populate cast targets
			ctargets := make(map[string]ability.CastTarget)

			if tt, ok := ef.Targets[key]; ok {
				if !tt.PositionTargetID.IsZero() {
					if t, ok := c.Targets[tt.PositionTargetID.String()]; ok {
						ctargets[key] = t
					}
				}
			}

			ev := event.E{
				ID:       ulid.NewID(),
				EntityID: t.ID,
				SourceID: sourceID,
				At:       c.At,
				Effect: ability.CastEffect{
					Effect:  ef,
					Targets: ctargets,
				},
			}

			result[ev.ID.String()] = ev
		}
	}

	return result, nil
}
