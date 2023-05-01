package main

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/rtc"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/panjf2000/ants/v2"
	"github.com/pion/webrtc/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	refreshRate   = 1000
	entitiesBatch = 50
	nRoutines     = 20
)

type antsLogger struct {
	zerolog.Logger
}

func (l *antsLogger) Printf(format string, args ...any) {
	l.Logger.Info().Msgf(format, args...)
}

func (h *handler) SendEntity(ctx context.Context, d *webrtc.DataChannel, pc entity.PC) error { //nolint: gocognit
	ctx, cancel := context.WithCancel(ctx)

	d.OnError(func(err error) {
		logger := log.With().Str("method", "send_entity").Str("pc", pc.ID.String()).Logger()

		logger.Error().Err(err).Msg("data channel error")

		cancel()
	})

	d.OnClose(func() {
		logger := log.With().Str("method", "send_entity").Str("pc", pc.ID.String()).Logger()

		logger.Info().Msg("data channel closed")

		cancel()
	})

	d.OnOpen(func() {
		logger := log.With().Str("method", "send_entity").Str("pc", pc.ID.String()).Logger()

		logger.Info().Msg("channel opened")

		p, err := ants.NewPool(
			nRoutines,
			ants.WithLogger(&antsLogger{logger}),
			ants.WithPreAlloc(true),
			ants.WithNonblocking(true),
		)
		if err != nil {
			logger.Error().Err(err).Msg("failed to create pools")

			return
		}

		logger = logger.With().Str("entity_id", pc.EntityID.String()).Logger()

		logger.Info().Msg("connected")

		t := time.NewTicker(refreshRate * time.Millisecond)

		for {
			select {
			case _ = <-ctx.Done():
				logger.Info().Err(ctx.Err()).Msg("done")

				// New context
				ctx := context.Background()

				logger = logger.With().
					Str("context", "clean_entity").
					Str("entity_id", pc.EntityID.String()).
					Logger()

				if err := h.entity.CreateBackupFromEntity(ctx, pc.EntityID); err != nil {
					logger.Error().Err(err).Msg("failed to create backup from entity")

					return
				}

				// #Delete connection
				if err := h.rtc.Delete(ctx, rtc.Filter{
					ID: pc.UserID,
				}); err != nil {
					logger.Error().Err(err).Msg("failed to delete rtc connection")

					return
				}

				logger.Info().Msg("done")

				return
			case _ = <-t.C:
				// Fetch current entity
				current, err := h.entity.Fetch(ctx,
					entity.Filter{
						ID: pc.EntityID,
					},
				)
				if err != nil {
					logger.Error().Err(err).Msg("failed to fetch current entity")

					continue
				}

				// Fetch current cell
				ce, err := h.room.FetchCell(ctx,
					room.FilterCell{
						ID: current.CellID,
					},
				)
				if err != nil {
					logger.Error().Err(err).Msg("failed to fetch current cell")

					continue
				}

				// Convert contiguous cell ids
				cellIDs := [9]ulid.ID{current.CellID}

				for n, id := range ce.Contiguous {
					cellIDs[n] = id
				}

				for _, id := range cellIDs {
					id := id

					if id.IsZero() {
						continue
					}

					p.Submit(func() {
						var state []byte

						for {
							entities, st, err := h.entity.FetchMany(ctx,
								entity.Filter{
									CellID: id,
									State:  state,
									Size:   entitiesBatch,
								},
							)
							if err != nil {
								logger.Error().Err(err).Msg("failed to fetch entites")

								return
							}

							state = st

							msg := &dto.ListEntityResp{
								Entities: entities,
								State:    state,
							}

							raw, err := msg.Marshal()
							if err != nil {
								logger.Error().Err(err).Msg("failed to marshal message")

								return
							}

							if err := d.Send(raw); err != nil {
								logger.Error().Err(err).Msg("failed to send message")

								return
							}

							if len(state) == 0 {
								logger.Info().Msg("success")

								return
							}
						}
					})
				}
			}
		}
	})

	return nil
}
