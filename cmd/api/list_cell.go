package main

import (
	"context"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/room/dto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListCell(ctx context.Context, req *dto.ListCellReq) (*dto.ListCellResp, error) {
	logger := log.With().Str("method", "list_cell").Logger()

	if req == nil {
		return &dto.ListCellResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListCellResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch pcs
	cells, _, err := h.room.FetchManyCell(ctx,
		room.FilterCell{
			IDs:  req.IDs,
			Size: len(req.IDs),
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pcs")

		return &dto.ListCellResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.ListCellResp{Cells: cells}, nil
}
