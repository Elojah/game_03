package main

import (
	"context"
	"errors"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/room/dto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateRoomUser(ctx context.Context, req *dto.CreateRoomUserReq) (*room.User, error) {
	logger := log.With().Str("method", "create_room_user").Logger()

	if req == nil {
		return &room.User{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &room.User{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch public room to ensure room exist and is public
	r, err := h.room.FetchPublic(ctx, room.FilterPublic{
		RoomID: req.RoomID,
	})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("room not found")

			return &room.User{}, status.New(codes.NotFound, err.Error()).Err()
		}
		logger.Error().Err(err).Msg("failed to fetch public room")

		return &room.User{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert room user
	ru := room.User{
		UserID: u.ID,
		RoomID: r.RoomID,
		Role:   int32(room.Member),
	}
	if err := h.room.InsertUser(ctx, ru); err != nil {
		logger.Error().Err(err).Msg("failed to create room user")

		return &room.User{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &ru, nil
}
