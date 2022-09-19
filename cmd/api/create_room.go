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
const (
	height, width         = 3, 3
	cellHeight, cellWidth = 1080, 1920
)

// var defaultTilemap = ulid.MustParse("01GDB4TQD4WS8XNEW9MP5YT1C5")

var defaultTilemap = ulid.MustParse("01FG1V2RER0529MJTZ9WE5HJF3")

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
	for i, cl := range cells {
		for j, c := range cl {
			c.Tilemap = defaultTilemap
			c.X = int64(j)
			c.Y = int64(i)

			if err := h.room.InsertCell(ctx, c); err != nil {
				logger.Error().Err(err).Msg("failed to create cell")

				return &room.R{}, status.New(codes.Internal, err.Error()).Err()
			}

			if err := h.room.InsertWorldCell(ctx, room.WorldCell{
				WorldID: w.ID,
				CellID:  c.ID,
				X:       c.X,
				Y:       c.Y,
			}); err != nil {
				logger.Error().Err(err).Msg("failed to create world cell")

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
