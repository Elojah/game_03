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

func (h *handler) ListWorld(ctx context.Context, req *dto.ListWorldReq) (*dto.ListWorldResp, error) {
	logger := log.With().Str("method", "list_world").Logger()

	if req == nil {
		return &dto.ListWorldResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListWorldResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	ws, state, err := h.room.FetchManyWorld(ctx, room.FilterWorld{
		IDs: req.IDs,
		All: req.All,

		Size:  int(req.Size_),
		State: req.State,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch world")

		return &dto.ListWorldResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	result := dto.ListWorldResp{Worlds: ws, State: state}

	logger.Info().Msg("success")

	return &result, nil
}
