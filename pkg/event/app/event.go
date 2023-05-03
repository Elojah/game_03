package app

import (
	"context"
	"strconv"

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

	entity entity.App
}

func (a App) Subscribe(ctx context.Context, entityID ulid.ID) error {
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
				lasts, err := a.entity.FetchManyCache(ctx, entity.FilterCache{
					ID:   entityID,
					Min:  "-inf",
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
				if err := a.entity.DeleteCache(ctx, entity.FilterCache{
					ID:   entityID,
					Min:  at,
					Max:  "+inf",
					Size: -1,
				}); err != nil {
					logger.Error().Err(err).Msg("failed to delete entities")

					continue
				}

				// Eval all events since at (should always include current one)
				events, err := a.FetchMany(ctx, event.Filter{
					EntityID: entityID,
					Min:      at,
					Max:      "+inf",
					Size:     -1,
				})
				if err != nil {
					logger.Error().Err(err).Msg("failed to fetch events")

					continue
				}

				for _, ev := range events {
					e = ev.Eval(e)

					if err := a.entity.InsertCache(ctx, e); err != nil {
						logger.Error().Err(err).Msg("failed to fetch events")

						continue
					}
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

		if err := a.Insert(ctx, ev); err != nil {
			log.Error().Err(err).Msg("failed to insert event")

			return
		}

		ch <- ev.At
	})
}
