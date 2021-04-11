package main

import (
	"context"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateRoom(ctx context.Context, req *room.R) (*room.R, error) {
	logger := log.With().Str("method", "create_room").Logger()

	if req == nil {
		return &room.R{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx)
	if err != nil {
		return &room.R{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	req.ID = ulid.NewID()

	// #Insert room
	if err := h.room.Insert(ctx, *req); err != nil {
		logger.Error().Err(err).Msg("failed to create room")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &room.R{}, nil
}
