package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) UpdatePCPreferences(ctx context.Context, req *entity.PCPreferences) (*entity.PCPreferences, error) {
	logger := log.With().Str("method", "update_pc_preferences").Logger()

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
		return &entity.PCPreferences{}, status.New(codes.PermissionDenied, err.Error()).Err()
	}

	// #Fetch pc preferences
	if err := h.entity.UpdatePCPreferences(
		ctx,
		entity.FilterPCPreferences{
			ID:   req.ID,
			PCID: req.PCID,
		},
		entity.PatchPCPreferences{
			AbilityHotbars: req.AbilityHotbars,
		},
	); err != nil {
		logger.Error().Err(err).Msg("failed to update pc preferences")

		return &entity.PCPreferences{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return req, nil
}
