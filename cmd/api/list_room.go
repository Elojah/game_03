package main

import (
	"context"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/room/dto"
	gulid "github.com/elojah/game_03/pkg/ulid"
	"github.com/elojah/game_03/pkg/user"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListRoom(ctx context.Context, req *dto.ListRoomReq) (*dto.ListRoomResp, error) {
	logger := log.With().Str("method", "list_room").Logger()

	if req == nil {
		return &dto.ListRoomResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		logger.Error().Err(err).Msg("failed to authenticate")

		return &dto.ListRoomResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch user rooms
	rus, state, err := h.room.FetchManyUser(ctx,
		room.FilterUser{
			UserID: u.ID,
			Size:   int(req.Size_),
			State:  req.State,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch room user")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	if len(rus) == 0 {
		logger.Info().Msg("success")

		return &dto.ListRoomResp{}, nil
	}

	// #Fetch rooms
	rooms, _, err := h.room.FetchMany(ctx, room.Filter{
		IDs:  room.Users(rus).RoomIDs(),
		Size: len(rus),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch rooms")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Populate rooms with owner
	ownerIDs := make([]gulid.ID, 0, len(rooms))

	for _, r := range rooms {
		ownerIDs = append(ownerIDs, r.OwnerID)
	}

	owners, _, err := h.user.FetchMany(ctx, user.Filter{
		IDs:  ownerIDs,
		Size: len(ownerIDs),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch owners")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	ownersMap := make(map[ulid.ULID]user.U, len(owners))
	for _, o := range owners {
		ownersMap[o.ID.ULID()] = o
	}

	// #Populate response
	result := dto.ListRoomResp{
		Rooms: make([]dto.Room, 0, len(rooms)),
		State: state,
	}

	for _, r := range rooms {
		result.Rooms = append(result.Rooms, dto.Room{
			Room:  r,
			Owner: ownersMap[r.OwnerID.ULID()],
		})
	}

	logger.Info().Msg("success")

	return &result, nil
}
