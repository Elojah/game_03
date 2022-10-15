package main

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) MigrateUp(ctx context.Context, req *types.StringValue) (*types.Empty, error) {
	logger := log.With().Str("method", "migrate_up").Logger()

	logger = logger.With().Str("dir", req.Value).Logger()

	logger.Info().Msg("start migration")

	if err := h.migrate.Up(ctx, req.Value); err != nil {
		logger.Error().Err(err).Msg("migration failed")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &types.Empty{}, nil
}
