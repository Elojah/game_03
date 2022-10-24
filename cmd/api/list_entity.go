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

func (h *handler) ListEntity(ctx context.Context, req *dto.ListEntityReq) (*dto.ListEntityResp, error) {
	logger := log.With().Str("method", "list_entity").Logger()

	if req == nil {
		return &dto.ListEntityResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.AuthSession(ctx)
	if err != nil {
		return &dto.ListEntityResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Check request
	if err := req.Check(); err != nil {
		return &dto.ListEntityResp{}, status.New(codes.InvalidArgument, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Fetch pcs
	entities, state, err := h.entity.FetchMany(ctx,
		entity.Filter{
			IDs:     req.IDs,
			CellIDs: req.CellIDs,
			State:   req.State,
			Size:    int(req.Size_),
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch entites")

		return &dto.ListEntityResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	result := dto.ListEntityResp{
		Entities: entities,
		State:    state,
	}

	logger.Info().Msg("success")

	return &result, nil
}
