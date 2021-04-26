package main

import (
	"errors"
	"time"

	"github.com/elojah/game_03/cmd/api/grpc"
	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ConnectPC(req *entity.PC, stream grpc.API_ConnectPCServer) error {
	ctx := stream.Context()
	logger := log.With().Str("method", "create_pc").Logger()

	if req == nil {
		return status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx)
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
		logger.Error().Err(err).Msg("failed to fetch pc connect")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch PC
	pc, err := h.entity.FetchPC(ctx, entity.FilterPC{
		ID: &req.ID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to create pc")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch Room
	r, err := h.room.Fetch(ctx, room.Filter{
		ID: &pc.RoomID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch room")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch world
	w, err := h.room.FetchWorld(ctx, room.FilterWorld{
		ID: &r.WorldID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch world")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch entity from PC
	e, err := h.entity.Fetch(ctx, entity.Filter{
		ID: &pc.EntityID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch cell from position
	x, y := w.Cell(e.X, e.Y)
	c, err := h.room.FetchCell(ctx, room.FilterCell{
		Keys: []room.KeyCell{{
			WorldID: w.ID,
			X:       x,
			Y:       y,
		}},
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return status.New(codes.Internal, err.Error()).Err()
	}

	if err := h.entity.Insert(ctx, entity.E{
		ID:     e.ID,
		CellID: c.ID,
		X:      e.X,
		Y:      e.Y,
		Rot:    e.Rot,
		Radius: e.Radius,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return status.New(codes.Internal, err.Error()).Err()
	}

	// #Create connect PC
	if err := h.entity.UpsertPCConnect(ctx, entity.PCConnect{
		ID:          req.ID,
		ConnectedAt: time.Now().Unix(),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create pc connect")

		return status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return nil
}
