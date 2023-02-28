package main

import (
	"context"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/room/dto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListRoomPublic(ctx context.Context, req *dto.ListRoomReq) (*dto.ListRoomResp, error) {
	logger := log.With().Str("method", "list_room_public").Logger()

	if req == nil {
		return &dto.ListRoomResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("failed to authenticate")

		return &dto.ListRoomResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	_ = u

	// #Fetch public rooms
	rooms, state, err := h.room.FetchManyPublic(ctx, room.FilterPublic{
		All:   true,
		State: req.State,
		Size:  int(req.Size_),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch public rooms")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Populate response
	result := dto.ListRoomResp{
		Rooms: make([]dto.Room, 0, len(rooms)),
		State: state,
	}

	for _, r := range rooms {
		result.Rooms = append(result.Rooms, dto.Room{
			Room: room.R{
				ID: r.RoomID,
			},
		})
	}

	logger.Info().Msg("success")

	return &result, nil
}
