package main

import (
	"context"
	"fmt"

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

	fmt.Println(req)

	// #Authenticate
	_, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListWorldResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if err := req.Check(); err != nil {
		logger.Error().Err(err).Msg("request check failed")

		return &dto.ListWorldResp{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	ws, _, err := h.room.FetchManyWorld(ctx, room.FilterWorld{
		IDs:  req.IDs,
		Size: len(req.IDs),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch world")

		return &dto.ListWorldResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	result := dto.ListWorldResp{Worlds: ws}

	logger.Info().Msg("success")

	return &result, nil
}
