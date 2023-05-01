package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog/log"
)

func (h *handler) ReceiveEntity(ctx context.Context, d *webrtc.DataChannel, pc entity.PC) error {
	d.OnMessage(func(msg webrtc.DataChannelMessage) {
		logger := log.With().Str("method", "receive_entity").Logger()

		var e entity.E

		if err := e.Unmarshal(msg.Data); err != nil {
			logger.Error().Err(err).Msg("failed to read entity")

			return
		}

		if !pc.EntityID.Equal(e.ID) {
			logger.Error().Msg("unauthorized entity update")

			return
		}

		if err := h.entity.Update(ctx, entity.Filter{
			ID: e.ID,
		}, entity.Patch{
			X:           &e.X,
			Y:           &e.Y,
			AnimationID: e.AnimationID,
			CellID:      e.CellID,
		}); err != nil {
			logger.Error().Err(err).Msg("failed to update entity")
		}

		logger.Info().Msg("success")

	})

	return nil
}
