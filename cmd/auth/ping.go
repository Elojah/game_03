package main

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
)

func (h *handler) Ping(ctx context.Context, req *types.Empty) (*types.Empty, error) {
	logger := log.With().Str("method", "ping").Logger()

	logger.Info().Msg("success")

	return &types.Empty{}, nil
}
