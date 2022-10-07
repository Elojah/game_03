package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateAnimation(ctx context.Context, req *dto.CreateAnimationReq) (*types.Empty, error) {
	logger := log.With().Str("method", "create_animation").Logger()

	if req == nil {
		return &types.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	// _, err := h.user.Auth(ctx)
	// if err != nil {
	// 	return &types.Empty{}, status.New(codes.Unauthenticated, err.Error()).Err()
	// }

	// if _, err := h.entity.Fetch(ctx, entity.Filter{
	// 	ID: &req.EntityID,
	// }); err != nil {
	// 	if errors.As(err, &gerrors.ErrNotFound{}) {
	// 		logger.Error().Err(err).Msg("entity not found")

	// 		return &types.Empty{}, status.New(codes.NotFound, err.Error()).Err()
	// 	}

	// 	logger.Error().Err(err).Msg("failed to fetch entity")

	// 	return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	// }

	// Parse ids
	id, err := ulid.Parse(req.ID)
	if err != nil {
		logger.Error().Err(err).Msg("invalid ulid")

		return &types.Empty{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	sheetID, err := ulid.Parse(req.SheetID)
	if err != nil {
		logger.Error().Err(err).Msg("invalid ulid")

		return &types.Empty{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	duplicateID, err := ulid.Parse(req.DuplicateID)
	if err != nil {
		logger.Error().Err(err).Msg("invalid ulid")

		return &types.Empty{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	req.Animation.ID = id
	req.Animation.SheetID = sheetID
	req.Animation.DuplicateID = duplicateID

	// Create animation
	if err := h.entity.InsertAnimation(ctx, req.Animation); err != nil {
		logger.Error().Err(err).Msg("failed to create animation")

		return &types.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &types.Empty{}, nil
}
