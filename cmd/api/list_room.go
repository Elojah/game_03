package main

import (
	"context"
	"fmt"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/room/dto"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/elojah/game_03/pkg/user"
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
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListRoomResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch pcs
	pcs, _, err := h.entity.FetchManyPC(ctx,
		entity.FilterPC{
			UserID: &ses.UserID,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pcs")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	if len(pcs) == 0 {
		logger.Info().Msg("success")

		return &dto.ListRoomResp{}, nil
	}

	// #Fetch rooms
	rooms, cursor, err := h.room.FetchMany(ctx, room.Filter{
		IDs:   entity.PCs(pcs).RoomIDs(),
		Size:  int(req.First),
		State: []byte(req.After),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch rooms")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	fmt.Println(rooms, cursor)
	func() {
		// #Fetch rooms
		rooms, cursor, err := h.room.FetchMany(ctx, room.Filter{
			IDs:   entity.PCs(pcs).RoomIDs(),
			Size:  1,
			State: cursor,
		})

		fmt.Println(rooms, cursor, err)
	}()

	// #Populate rooms with owner
	ownerIDs := make([]ulid.ID, 0, len(rooms))

	for _, r := range rooms {
		ownerIDs = append(ownerIDs, r.OwnerID)
	}

	owners, _, err := h.user.FetchMany(ctx, user.Filter{
		IDs: ownerIDs,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch owners")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	ownersMap := make(map[ulid.ID]user.U, len(owners))
	for _, o := range owners {
		ownersMap[o.ID] = o
	}

	// #Populate response
	result := dto.ListRoomResp{
		Rooms: make([]dto.Room, 0, len(rooms)),
	}

	for _, r := range rooms {
		result.Rooms = append(result.Rooms, dto.Room{
			Room:  r,
			Owner: ownersMap[r.OwnerID],
		})
	}

	logger.Info().Msg("success")

	return &result, nil
}
