package main

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/user"
	"github.com/elojah/game_03/pkg/user/dto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateSession(ctx context.Context, req *dto.CreateSessionReq) (*dto.CreateSessionResp, error) {
	logger := log.With().Str("method", "create_session").Logger()

	if req == nil {
		logger.Error().Err(gerrors.ErrNullRequest{}).Msg("null request")

		return &dto.CreateSessionResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.CreateSessionResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch PC
	pc, err := h.entity.FetchPC(ctx, entity.FilterPC{
		ID:      req.PCID,
		WorldID: req.WorldID,

		UserID: u.ID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pc")

		return &dto.CreateSessionResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	ses := user.Session{
		ID:     pc.ID,
		UserID: u.ID,
		At:     time.Now().Unix(),
	}

	if err := h.user.InsertSession(ctx, ses); err != nil {
		logger.Error().Err(err).Msg("failed to create session")

		return &dto.CreateSessionResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	raw, err := ses.Encrypt()
	if err != nil {
		logger.Error().Err(err).Msg("failed to encrypt session")

		return &dto.CreateSessionResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.CreateSessionResp{Token: raw}, nil
}
