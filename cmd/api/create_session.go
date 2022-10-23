package main

import (
	"context"
	"errors"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/user"
	"github.com/elojah/game_03/pkg/user/dto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateSession(ctx context.Context, req *dto.CreateSessionReq) (*user.Session, error) {
	logger := log.With().Str("method", "get_session").Logger()

	if req == nil {
		logger.Error().Err(gerrors.ErrNullRequest{}).Msg("null request")

		return &user.Session{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx)
	if err != nil {
		return &user.Session{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch room
	ro, err := h.room.Fetch(ctx, room.Filter{
		ID: req.RoomID,
	})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("missing room")

			return &user.Session{}, status.New(codes.FailedPrecondition, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch room user")

		return &user.Session{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch PC
	pc, err := h.entity.FetchPC(ctx, entity.FilterPC{
		ID:      req.PCID,
		WorldID: ro.WorldID,

		UserID: u.ID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pc")

		return &user.Session{}, status.New(codes.Internal, err.Error()).Err()
	}

	ses := user.Session{
		ID:     pc.ID,
		UserID: u.ID,
		At:     time.Now().Unix(),
	}

	if err := h.user.InsertSession(ctx, ses); err != nil {
		logger.Error().Err(err).Msg("failed to create session")

		return &user.Session{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &ses, nil
}
