package main

import (
	"context"
	"errors"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetPCPreferences(ctx context.Context, req *entity.PC) (*entity.PCPreferences, error) {
	logger := log.With().Str("method", "get_pc_preferences").Logger()

	if req == nil {
		return &entity.PCPreferences{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &entity.PCPreferences{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	if _, err := h.entity.FetchPC(ctx, entity.FilterPC{
		ID:     req.ID,
		UserID: u.ID,
	}); err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("pc not found")

			return &entity.PCPreferences{}, status.New(codes.NotFound, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch template")

		return &entity.PCPreferences{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch pc preferences
	pcp, err := h.entity.FetchPCPreferences(ctx,
		entity.FilterPCPreferences{
			PCID: req.ID,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pc preferences")

		return &entity.PCPreferences{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &pcp, nil
}
