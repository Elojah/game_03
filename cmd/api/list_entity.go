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

	// #Check request
	if err := req.Check(); err != nil {
		return &dto.ListEntityResp{}, status.New(codes.InvalidArgument, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListEntityResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch pcs
	entities, _, err := h.entity.FetchMany(ctx,
		entity.Filter{
			IDs:     req.IDs,
			CellIDs: req.CellIDs,
			State:   req.State,
			Size:    len(req.IDs),
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch entites")

		return &dto.ListEntityResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Filter out non user entities
	var result dto.ListEntityResp

	for _, e := range entities {
		if e.UserID == ses.UserID {
			result.Entities = append(result.Entities, e)
		}
	}

	logger.Info().Msg("success")

	return &result, nil
}
