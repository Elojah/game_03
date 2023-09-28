package main

import (
	"context"
	"errors"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/pbtypes"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) CreateAnimation(ctx context.Context, req *dto.CreateAnimationReq) (*pbtypes.Empty, error) {
	logger := log.With().Str("method", "create_animation").Logger()

	if req == nil {
		return &pbtypes.Empty{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	// _, err := h.user.Auth(ctx, "access")
	// if err != nil {
	// 	return &pbtypes.Empty{}, status.New(codes.Unauthenticated, err.Error()).Err()
	// }

	// entity.id has priority over entity.template name
	if !req.Animation.EntityID.IsZero() {
		if _, err := h.entity.Fetch(ctx, entity.Filter{
			ID: req.Animation.EntityID,
		}); err != nil {
			if errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("entity not found")

				return &pbtypes.Empty{}, status.New(codes.NotFound, err.Error()).Err()
			}

			logger.Error().Err(err).Msg("failed to fetch entity")

			return &pbtypes.Empty{}, status.New(codes.Internal, err.Error()).Err()
		}
	} else if req.EntityTemplate != "" {
		template, err := h.entity.FetchTemplate(ctx, entity.FilterTemplate{
			Name: &req.EntityTemplate,
		})
		if err != nil {
			if errors.As(err, &gerrors.ErrNotFound{}) {
				logger.Error().Err(err).Msg("template not found")

				return &pbtypes.Empty{}, status.New(codes.NotFound, err.Error()).Err()
			}

			logger.Error().Err(err).Msg("failed to fetch template")

			return &pbtypes.Empty{}, status.New(codes.Internal, err.Error()).Err()
		}

		req.Animation.EntityID = template.EntityID
	} else {
		err := gerrors.ErrMissingEntity{AnimationID: req.Animation.ID.String()}

		logger.Error().Err(err).Msg("missing entity")

		return &pbtypes.Empty{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	if req.SheetID.IsZero() {
		err := gerrors.ErrMissingSheet{AnimationID: req.Animation.ID.String()}

		logger.Error().Err(err).Msg("missing sheet")

		return &pbtypes.Empty{}, status.New(codes.InvalidArgument, err.Error()).Err()
	}

	// populate with default ids
	if req.ID.IsZero() {
		req.ID = ulid.NewID()
	}

	if req.DuplicateID.IsZero() {
		req.DuplicateID = ulid.NewID()
	}

	// Create animation
	if err := h.entity.InsertAnimation(ctx, req.Animation); err != nil {
		logger.Error().Err(err).Msg("failed to create animation")

		return &pbtypes.Empty{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &pbtypes.Empty{}, nil
}
