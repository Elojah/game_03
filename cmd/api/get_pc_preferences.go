package main

import (
	"context"

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
	session, err := h.user.AuthSession(ctx)
	if err != nil {
		return &entity.PCPreferences{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// Check PC ID is session ID
	if session.ID.Compare(req.ID) != 0 {
		return &entity.PCPreferences{}, status.New(codes.PermissionDenied, gerrors.ErrPermissionForbidden{}.Error()).Err()
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
