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

func (h *handler) ListTemplate(ctx context.Context, req *dto.ListTemplateReq) (*dto.ListTemplateResp, error) {
	logger := log.With().Str("method", "list_template").Logger()

	if req == nil {
		return &dto.ListTemplateResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.ListTemplateResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	ts, state, err := h.entity.FetchManyTemplate(ctx, entity.FilterTemplate{
		All: req.All,

		Size:  int(req.Size_),
		State: req.State,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch template")

		return &dto.ListTemplateResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	result := dto.ListTemplateResp{Templates: ts, State: state}

	logger.Info().Msg("success")

	return &result, nil
}
