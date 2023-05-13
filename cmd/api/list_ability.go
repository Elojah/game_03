package main

import (
	"context"

	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/game_03/pkg/ability/dto"
	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListAbility(ctx context.Context, req *dto.ListAbilityReq) (*dto.ListAbilityResp, error) {
	logger := log.With().Str("method", "list_ability").Logger()

	if req == nil {
		return &dto.ListAbilityResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	_, err := h.user.AuthSession(ctx)
	if err != nil {
		return &dto.ListAbilityResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch abilities
	eabs, state, err := h.entity.FetchManyAbility(ctx,
		entity.FilterAbility{
			EntityID: req.EntityID,
			State:    req.State,
			Size:     int(req.Size_),
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch entity abilities")

		return &dto.ListAbilityResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	abilityIDs := func(eabs []entity.Ability) []ulid.ID {
		result := make([]ulid.ID, len(eabs))

		for i, eab := range eabs {
			result[i] = eab.AbilityID
		}

		return result
	}(eabs)

	abs, _, err := h.ability.FetchMany(ctx, ability.Filter{
		IDs:  abilityIDs,
		Size: len(abilityIDs),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch abilities")

		return &dto.ListAbilityResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &dto.ListAbilityResp{
		Abilities: abs,
		State:     state,
	}, nil
}
