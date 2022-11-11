package main

import (
	"context"
	"encoding/xml"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/tile"
	"github.com/elojah/game_03/pkg/tile/dto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateTileset(ctx context.Context, req *dto.CreateTilesetReq) (*dto.CreateTilesetResp, error) {
	logger := log.With().Str("method", "create_tileset").Logger()

	if req == nil {
		return &dto.CreateTilesetResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	var ts tile.Set

	if err := xml.Unmarshal(req.Set, &ts); err != nil {
		logger.Error().Err(err).Msg("failed to unmarshal set")

		return &dto.CreateTilesetResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	ts.ID = req.ID

	if err := h.tile.InsertSet(ctx, ts); err != nil {
		logger.Error().Err(err).Msg("failed to create tileset")

		return &dto.CreateTilesetResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.CreateTilesetResp{
		ID: ts.ID,
	}, nil
}
