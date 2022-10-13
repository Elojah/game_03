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

// TMP DATA FOR DEV WIP.

var defaultWorldID = ulid.MustParse("01GF89A3E81XM00XD46JB4MRBP")

// !TMP DATA FOR DEV WIP

func (h *handler) CreateRoom(ctx context.Context, req *room.R) (*room.R, error) {
	logger := log.With().Str("method", "create_room").Logger()

	if req == nil {
		return &room.R{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return &room.R{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Set new room values
	req.ID = ulid.NewID()
	req.WorldID = defaultWorldID
	req.OwnerID = ses.UserID

	// #Insert room
	if err := h.room.Insert(ctx, *req); err != nil {
		logger.Error().Err(err).Msg("failed to create room")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert room user
	if err := h.room.InsertUser(ctx, room.User{
		UserID: ses.UserID,
		RoomID: req.ID,
		Role:   int32(room.Owner),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create room")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return req, nil
}
