package main

import (
	"context"
	"errors"

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
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &room.R{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	worldID, err := h.room.CopyWorld(ctx, req.WorldID)
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("world not found")

			return &room.R{}, status.New(codes.NotFound, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to copy world")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Set new room values
	req.ID = ulid.NewID()
	req.OwnerID = u.ID
	req.WorldID = worldID

	// #Insert room
	if err := h.room.Insert(ctx, *req); err != nil {
		logger.Error().Err(err).Msg("failed to create room")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert room user
	if err := h.room.InsertUser(ctx, room.User{
		UserID: u.ID,
		RoomID: req.ID,
		Role:   int32(room.Owner),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create room")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert public room
	// TODO: use parameter to set public or not
	if err := h.room.InsertPublic(ctx, room.Public{
		ID:     ulid.NewID(),
		RoomID: req.ID,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create public room")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return req, nil
}
