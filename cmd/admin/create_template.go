package main

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateTemplate(ctx context.Context, req *dto.CreateTemplateReq) (*entity.Template, error) {
	logger := log.With().Str("method", "create_template").Logger()

	if req == nil {
		return &entity.Template{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	// _, err := h.user.Auth(ctx, "access")
	// if err != nil {
	// 	return &entity.Template{}, status.New(codes.Unauthenticated, err.Error()).Err()
	// }

	// Create entity
	e := entity.E{
		ID: ulid.NewID(),

		Name:    req.Name,
		Radius:  req.Entity.Radius,
		Objects: req.Entity.Objects,

		UserID: nil, // explicitly set nil
		CellID: nil, // explicitly set nil
	}
	if err := h.entity.Insert(ctx, e); err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return &entity.Template{}, status.New(codes.Internal, err.Error()).Err()
	}

	t := entity.Template{
		ID:       ulid.NewID(),
		EntityID: e.ID,
		Name:     req.Name,
	}

	// Create template
	if err := h.entity.InsertTemplate(ctx, t); err != nil {
		logger.Error().Err(err).Msg("failed to create template")

		return &entity.Template{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &t, nil
}
