package main

import (
	"errors"
	"time"

	"github.com/elojah/game_03/cmd/api/grpc"
	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	refreshRate = 100 * time.Millisecond
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

	for {
		select {
		case _ = <-ctx.Done():
			logger.Info().Msg("done")

			return ctx.Err()
		case _ = <-t.C:
			// TODO: fetch regions + fetch entities and send back into stream
			_ = 0
		}
	}
}
