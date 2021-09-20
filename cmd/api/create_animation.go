package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateAnimation(ctx context.Context, req *entity.Animation) (*types.Empty, error) {
	logger := log.With().Str("method", "create_animation").Logger()

	if req == nil {
		return &types.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx)
	if err != nil {
		return &types.Empty{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// // Create animation
	// if err := h.entity.InsertAnimation(ctx, *req); err != nil {
	// 	logger.Error().Err(err).Msg("failed to create animation")

	// 	return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	// }

	logger.Info().Msg("success")

	return &types.Empty{}, nil
}
