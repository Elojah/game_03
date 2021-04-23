package main

import (
	"context"

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

var tileID = ulid.MustParse("01F3538E0FEXZ563X1NG24SHFN")

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
		// Tileset: map[int32]ulid.ID{
		// 	0: tileID,
		// },
	}
	if err := h.room.InsertWorld(ctx, w); err != nil {
		logger.Error().Err(err).Msg("failed to create world")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Create world cells
	tm := func() map[int64]ulid.ID {
		result := make(map[int64]ulid.ID, cellHeight*cellWidth)

		for i := int64(0); i < cellHeight; i++ {
			for j := int64(0); j < cellWidth; j++ {
				result[(i*cellWidth)+j] = tileID
			}
		}

		return result
	}()

	// TODO: add goroutine pool
	for i := int64(0); i < height/cellHeight; i++ {
		for j := int64(0); j < width/cellWidth; j++ {
			if err := h.room.InsertCell(ctx, room.Cell{
				WorldID: w.ID,
				X:       i,
				Y:       j,
				Tilemap: tm,
			}); err != nil {
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

	// #Insert owner PC
	if err := h.entity.InsertPC(ctx, entity.PC{
		ID:     ulid.NewID(),
		UserID: ses.UserID,
		RoomID: req.ID,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create pc")

		return &room.R{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &room.R{}, nil
}
