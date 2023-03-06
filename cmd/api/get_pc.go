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

func (h *handler) GetPC(ctx context.Context, req *dto.GetPCReq) (*dto.PC, error) {
	logger := log.With().Str("method", "get_pc").Logger()

	if req == nil {
		return &dto.PC{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.PC{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch pc
	pc, err := h.entity.FetchPC(ctx,
		entity.FilterPC{
			ID:      req.ID,
			WorldID: req.WorldID,
			UserID:  u.ID,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pc")

		return &dto.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch entity
	e, err := h.entity.FetchBackup(ctx, entity.Filter{
		ID: pc.EntityID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch entity")

		return &dto.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.PC{
		PC:     pc,
		Entity: e,
	}, nil
}
