package main

import (
	"context"
	"errors"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var defaultSpriteID = ulid.MustParse("01F5X76JW6W6PS824SFXC9R8BB")

func (h *handler) CreatePC(ctx context.Context, req *dto.CreatePCReq) (*entity.PC, error) {
	logger := log.With().Str("method", "create_pc").Logger()

	if req == nil {
		return &entity.PC{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return &entity.PC{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Check if user is room user
	if _, err := h.room.FetchUser(ctx, room.FilterUser{
		UserID: &ses.UserID,
		RoomID: &req.RoomID,
	}); err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("missing room user")

			return &entity.PC{}, status.New(codes.FailedPrecondition, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch room user")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch room
	ro, err := h.room.Fetch(ctx, room.Filter{
		ID: &req.RoomID,
	})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("missing room")

			return &entity.PC{}, status.New(codes.FailedPrecondition, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch room user")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch 0,0 cell
	c, err := h.room.FetchWorldCell(ctx, room.FilterWorldCell{
		WorldID: ro.WorldID,
		X:       0,
		Y:       0,
	})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("missing cell")

			return &entity.PC{}, status.New(codes.FailedPrecondition, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch world cell")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert entity backup
	bu := entity.Backup{
		ID:     ulid.NewID(),
		UserID: ses.UserID,
		CellID: c.CellID,
		X:      0,
		Y:      0,
		Rot:    0,
		Radius: 10, // nolint: gomnd
		At:     time.Now().UnixNano(),
	}
	if err := h.entity.InsertBackup(ctx, bu); err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert owner PC
	if err := h.entity.InsertPC(ctx, entity.PC{
		ID:       ulid.NewID(),
		EntityID: bu.ID,
		UserID:   ses.UserID,
		WorldID:  ro.WorldID,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create pc")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &entity.PC{}, nil
}
