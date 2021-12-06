package main

import (
	"io"

	ggrpc "github.com/elojah/game_03/cmd/api/grpc"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdateEntity(stream ggrpc.API_UpdateEntityServer) error {
	ctx := stream.Context()
	logger := log.With().Str("method", "update_entity").Logger()

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return status.New(codes.Unauthenticated, err.Error()).Err()
	}

	logger.Info().Msg("success")

	for {
		// Check context done
		select {
		case _ = <-ctx.Done():
			logger.Info().Msg("done")

			return ctx.Err()
		default:
		}

		// Receive entity
		e, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&types.Empty{})
		}

		if ses.UserID != e.UserID {
			err := errors.ErrInvalidCredentials{}
			logger.Error().Err(err).Msg("invalid sender")

			return status.New(codes.PermissionDenied, err.Error()).Err()
		}

		if err != nil {
			logger.Error().Err(err).Msg("failed to receive entity")

			return status.New(codes.Internal, err.Error()).Err()
		}

		if e == nil {
			return status.New(codes.InvalidArgument, errors.ErrNullRequest{}.Error()).Err()
		}

		logger = logger.With().Str("entity_id", e.ID.String()).Logger()

		if err := h.entity.Insert(ctx, *e); err != nil {
			logger.Error().Err(err).Msg("failed to insert entity")

			return status.New(codes.Internal, err.Error()).Err()
		}

		logger.Info().Msg("success")
	}
}
