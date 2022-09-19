package main

import (
	"context"

	"github.com/elojah/game_03/pkg/tile/dto"
	"github.com/rs/zerolog/log"
)

// TMP TEST.
func (h *handler) CreateTileset(ctx context.Context, req *dto.CreateTilesetReq) (*dto.CreateTilesetResp, error) {
	logger := log.With().Str("method", "create_tileset").Logger()

	logger.Info().Msg("success")

	return nil, nil
}
