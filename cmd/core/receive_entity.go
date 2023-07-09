package main

import (
	"context"
	"strconv"
	"time"

	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/game_03/pkg/entity"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
)

const (
	cleanEntities = 5000
)

func (h *handler) ReceiveEntity(ctx context.Context, d *webrtc.DataChannel, pc entity.PC) error {
	d.OnOpen(func() {
		logger := log.With().Str("method", "receive_entity").Str("subscribe", pc.EntityID.String()).Logger()
		ticker := time.NewTicker(cleanEntities * time.Millisecond)

		// entity cache cleaning
		go func() {
			for {
				select {
				case _ = <-ctx.Done():
					logger.Error().Err(ctx.Err()).Msg("context done")
					return
				case t := <-ticker.C:
					at := strconv.FormatInt(t.Add(-cleanEntities*time.Millisecond).UnixMilli(), 10)

					if err := h.entity.DeleteCache(ctx, entity.FilterCache{Max: at, Min: "-inf"}); err != nil {
						logger.Error().Err(err).Msg("entity delete cache failed")

						break
					}
				}
			}
		}()

		// blocking call
		if err := h.event.Eval(ctx, pc.EntityID); err != nil {
			logger.Error().Err(err).Msg("subscribe failed")

			return
		}
	})

	d.OnMessage(func(msg webrtc.DataChannelMessage) {
		logger := log.With().Str("method", "receive_entity").Logger()

		var c ability.Cast

		if err := c.Unmarshal(msg.Data); err != nil {
			logger.Error().Err(err).Msg("failed to read message")

			return
		}

		logger.Info().Int64("at", c.At).Msg("event received")

		events, err := h.event.CreateFromCast(ctx, pc.EntityID, c)
		if err != nil {
			logger.Error().Err(err).Msg("failed to create events from cast")

			return
		}

		for _, e := range events {
			if err := h.event.Publish(ctx, e); err != nil {
				logger.Error().Err(err).Msg("failed to publish event")

				return
			}
		}

		logger.Info().Msg("success")
	})

	return nil
}
