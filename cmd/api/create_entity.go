package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateEntity(ctx context.Context, req *entity.E) (*entity.E, error) {
	logger := log.With().Str("method", "create_entity").Logger()

	if req == nil {
		return &entity.E{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx)
	if err != nil {
		return &entity.E{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &entity.E{}, nil
}
