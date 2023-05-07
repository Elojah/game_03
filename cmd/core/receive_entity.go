package main

import (
	"context"

	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/game_03/pkg/entity"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
)

func (h *handler) ReceiveEntity(ctx context.Context, d *webrtc.DataChannel, pc entity.PC) error {
	d.OnOpen(func() {
		// blocking call
		if err := h.event.Eval(ctx, pc.EntityID); err != nil {
			logger := log.With().Str("method", "receive_entity").Str("subscribe", pc.EntityID.String()).Logger()

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

		// TODO: insert events here and publish only one last ping to all concerned entities

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
