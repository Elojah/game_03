package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListAnimation(ctx context.Context, req *dto.ListAnimationReq) (*dto.ListAnimationResp, error) {
	logger := log.With().Str("method", "list_animation").Logger()

	if req == nil {
		return &dto.ListAnimationResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListAnimationResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch animations
	ans, state, err := h.entity.FetchManyAnimation(ctx,
		entity.FilterAnimation{
			IDs:       req.IDs,
			EntityIDs: req.EntityIDs,
			State:     req.State,
			Size:      int(req.Size_),
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch animations")

		return &dto.ListAnimationResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.ListAnimationResp{
		Animations: ans,
		State:      state,
	}, nil
}
