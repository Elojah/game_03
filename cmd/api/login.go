package main

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
	"github.com/elojah/game_03/pkg/user"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Login(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	logger := log.With().Str("method", "login").Logger()

	// Create new session
	id := ulid.NewID()

	if err := h.user.UpsertSession(ctx, user.Session{ID: id}); err != nil {
		logger.Error().Err(err).Msg("failed to create new session")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &types.Empty{}, nil
}
