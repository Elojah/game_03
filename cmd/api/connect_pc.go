package main

import (
	"context"
	"errors"
	"time"

	"github.com/elojah/game_03/cmd/api/grpc"
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	refreshRate   = 100 * time.Millisecond
	entitiesBatch = 20
)

// nolint: gocognit
func (h *handler) ConnectPC(req *entity.PC, stream grpc.API_ConnectPCServer) error {
	ctx := stream.Context()
	logger := log.With().Str("method", "connect_pc").Logger()

	if req == nil {
		return status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Check if connect PC already exists
	if _, err := h.entity.FetchPCConnect(ctx, entity.FilterPCConnect{
		ID: &req.ID,
	}); err != nil {
		if !errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("pc already connected")

			return status.New(codes.FailedPrecondition, err.Error()).Err()
		}
	} else {
		err := gerrors.ErrPCAlreadyConnected{ID: req.ID.String()}

		logger.Error().Err(err).Msg("failed to fetch pc connect")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch PC
	pc, err := h.entity.FetchPC(ctx, entity.FilterPC{
		ID:      &req.ID,
		WorldID: &req.WorldID,

		UserID: &ses.UserID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pc")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch entity backup from PC
	e, err := h.entity.FetchBackup(ctx, entity.FilterBackup{
		ID: &pc.EntityID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// Insert real entity
	e.At = time.Now().UnixNano()
	if err := h.entity.Insert(ctx, e); err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// defer clean
	defer func(id ulid.ID) {
		// New context
		ctx := context.Background()
		logger = logger.With().
			Str("defer", "clean_entity").
			Str("entity_id", id.String()).
			Logger()

		// #Fetch entity
		e, err := h.entity.Fetch(ctx, entity.Filter{
			ID: &id,
		})
		if err != nil {
			logger.Error().Err(err).Msg("failed to fetch entity")

			return
		}

		// Insert backup entity
		e.At = time.Now().UnixNano()
		if err := h.entity.UpdateBackup(ctx, entity.Filter{
			ID: &id,
		}, e); err != nil {
			logger.Error().Err(err).Msg("failed to update backup")

			return
		}

		// #Delete entity
		if err := h.entity.Delete(ctx, entity.Filter{
			ID: &id,
		}); err != nil {
			logger.Error().Err(err).Msg("failed to delete entity")

			return
		}

		logger.Info().Msg("done")
	}(pc.EntityID)

	logger = logger.With().Str("entity_id", e.ID.String()).Logger()

	// #Create connect PC
	if err := h.entity.UpsertPCConnect(ctx, entity.PCConnect{
		ID:          req.ID,
		ConnectedAt: time.Now().Unix(),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create pc connect")

		return status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("connected")

	t := time.NewTicker(refreshRate)

	// TODO: remove debug
	staticDebugID := ulid.NewID()

	for {
		select {
		case _ = <-ctx.Done():
			logger.Info().Msg("done")

			return ctx.Err()
		case _ = <-t.C:
			// Fetch current entity
			current, err := h.entity.Fetch(ctx,
				entity.Filter{
					ID: &e.ID,
				},
			)
			if err != nil {
				logger.Error().Err(err).Msg("failed to fetch current entity")

				continue
			}

			// Fetch current cell
			ce, err := h.room.FetchCell(ctx,
				room.FilterCell{
					ID: &current.CellID,
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

			// #Fetch all entities
			var entities []entity.E

			var state []byte

			for {
				entities, state, err = h.entity.FetchMany(ctx,
					entity.Filter{
						CellIDs: cellIDs[:],
						State:   state,
						Size:    entitiesBatch,
					},
				)
				if err != nil {
					logger.Error().Err(err).Msg("failed to fetch entites")

					break
				}

				// TODO: remove debug
				entities = append(entities, entity.E{
					ID:          staticDebugID,
					UserID:      current.UserID,
					CellID:      current.CellID,
					Name:        "test debug pet",
					X:           current.X + 5, // nolint: gomnd
					Y:           current.Y + 5, // nolint: gomnd
					AnimationID: current.AnimationID,
				})

				if err := stream.SendMsg(&dto.ListEntityResp{
					Entities: entities,
					State:    state,
				}); err != nil {
					logger.Error().Err(err).Msg("failed to send message")

					break
				}

				if len(state) == 0 {
					logger.Info().Msg("success")

					break
				}
			}
		}
	}
}
