package main

import (
	"context"

	"github.com/elojah/game_03/pkg/pbtypes"
	"github.com/rs/zerolog/log"
)

func (h *handler) Ping(ctx context.Context, req *pbtypes.Empty) (*pbtypes.Empty, error) {
	logger := log.With().Str("method", "ping").Logger()

	logger.Info().Msg("success")

	return &pbtypes.Empty{}, nil
}
