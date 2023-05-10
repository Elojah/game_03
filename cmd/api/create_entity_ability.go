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

func (h *handler) CreateEntityAbility(ctx context.Context, req *dto.CreateEntityAbilityReq) (*dto.CreateEntityAbilityResp, error) {
	logger := log.With().Str("method", "create_entity_ability").Logger()

	if req == nil {
		return &dto.CreateEntityAbilityResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &dto.CreateEntityAbilityResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	req.Ability.ID = ulid.NewID()

	if err := h.ability.Insert(ctx, req.Ability); err != nil {
		logger.Error().Err(err).Msg("failed to create ability")

		return &dto.CreateEntityAbilityResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	if err := h.entity.InsertAbility(ctx, entity.Ability{
		AbilityID: req.Ability.ID,
		EntityID:  req.EntityID,
		LastCast:  0,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create entity ability")

		return &dto.CreateEntityAbilityResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.CreateEntityAbilityResp{
		AbilityID: req.Ability.ID,
		EntityID:  req.EntityID,
	}, nil
}
