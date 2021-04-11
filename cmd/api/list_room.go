package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/room/dto"
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

	// #Fetch user
	u, err := h.user.Fetch(ctx, user.Filter{ID: &ses.UserID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch user")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch pcs
	pcs, err := h.entity.FetchManyPC(ctx,
		entity.FilterPC{
			UserID: &u.ID,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pcs")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch pcs
	rooms, err := h.room.FetchMany(ctx, room.Filter{
		IDs: entity.PCs(pcs).RoomIDs(),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch rooms")

		return &dto.ListRoomResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.ListRoomResp{Rooms: rooms}, nil
}
