package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListPC(ctx context.Context, req *dto.ListPCReq) (*dto.ListPCResp, error) {
	logger := log.With().Str("method", "list_room").Logger()

	if req == nil {
		return &dto.ListPCResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListPCResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch pcs
	pcs, state, err := h.entity.FetchManyPC(ctx,
		entity.FilterPC{
			UserID: &ses.UserID,
			Size:   int(req.Size_),
			State:  req.State,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pcs")

		return &dto.ListPCResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.ListPCResp{PCs: pcs, State: state}, nil
}
