package main

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TMP DATA FOR DEV WIP

const (
	height, width         = 1000, 1000
	cellHeight, cellWidth = 200, 200
)

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

	// #Create world
	w := room.World{
		ID:         ulid.NewID(),
		Height:     height,
		Width:      width,
		CellHeight: cellHeight,
		CellWidth:  cellWidth,
	}
	if err := h.room.InsertWorld(ctx, w); err != nil {
		logger.Error().Err(err).Msg("failed to create world")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Create world cells
	cells := w.NewCells()

	// TODO: add goroutine pool
	for _, cl := range cells {
		for _, c := range cl {
			if err := h.room.InsertCell(ctx, c); err != nil {
				logger.Error().Err(err).Msg("failed to create cell")

				return &room.R{}, status.New(codes.Internal, err.Error()).Err()
			}
		}
	}

	// #Set new room values
	req.ID = ulid.NewID()
	req.WorldID = w.ID
	req.OwnerID = ses.UserID

	// #Insert room
	if err := h.room.Insert(ctx, *req); err != nil {
		logger.Error().Err(err).Msg("failed to create room")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert entity backup
	// TODO: define spawns for entity
	bu := entity.Backup{
		ID: ulid.NewID(),
		At: time.Now().UnixNano(),
	}
	if err := h.entity.InsertBackup(ctx, bu); err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert owner PC
	if err := h.entity.InsertPC(ctx, entity.PC{
		ID:       ulid.NewID(),
		EntityID: bu.ID,
		UserID:   ses.UserID,
		RoomID:   req.ID,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create pc")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &room.R{}, nil
}
