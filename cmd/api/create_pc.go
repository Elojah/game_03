package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/faction"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	maxAnimations = 100
)

func (h *handler) CreatePC(ctx context.Context, req *dto.CreatePCReq) (*entity.PC, error) {
	logger := log.With().Str("method", "create_pc").Logger()

	if req == nil {
		return &entity.PC{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	u, err := h.user.Auth(ctx, "access")
	if err != nil {
		return &entity.PC{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Check if user is room user
	if _, err := h.room.FetchUser(ctx, room.FilterUser{
		UserID: u.ID,
		RoomID: req.RoomID,
	}); err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("missing room user")

			return &entity.PC{}, status.New(codes.FailedPrecondition, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch room user")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch room
	ro, err := h.room.Fetch(ctx, room.Filter{
		ID: req.RoomID,
	})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("missing room")

			return &entity.PC{}, status.New(codes.FailedPrecondition, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch room user")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch 0,0 cell
	c, err := h.room.FetchWorldCell(ctx, room.FilterWorldCell{
		WorldID: ro.WorldID,
		X:       0,
		Y:       0,
	})
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("missing cell")

			return &entity.PC{}, status.New(codes.FailedPrecondition, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch world cell")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Create entity from template
	template, err := h.entity.FetchTemplate(ctx, entity.FilterTemplate{
		Name: &req.EntityTemplate,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch template")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	eTemplate, err := h.entity.Fetch(ctx, entity.Filter{
		ID: template.EntityID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch entity")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Create entity animations
	ans, _, err := h.entity.FetchManyAnimation(ctx, entity.FilterAnimation{
		EntityID: template.EntityID,
		Size:     maxAnimations,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch many animation")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	if len(ans) == 0 {
		err := gerrors.ErrMissingDefaultAnimations{EntityID: template.EntityID.String()}

		logger.Error().Err(err).Msg("missing animation")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	entityID := ulid.NewID()

	var idle ulid.ID

	for _, an := range ans {
		an.ID = ulid.NewID()
		an.EntityID = entityID

		// default animation name is always "idle"
		if an.Name == "idle" {
			idle = an.ID
		}

		if err := h.entity.InsertAnimation(ctx, an); err != nil {
			logger.Error().Err(err).Msg("failed to create animation")

			return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
		}
	}

	// #Create entity spawn

	// Fetch first spawn in world
	// TODO: pick algo + pick(create ?) ownerID
	sp, err := h.room.FetchWorldSpawn(ctx, room.FilterWorldSpawn{WorldID: ro.WorldID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch world spawn")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	if err := h.entity.InsertSpawn(ctx, entity.Spawn{
		EntityID: entityID,
		SpawnID:  sp.ID,
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create entity spawn")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert faction pc
	// TODO: add as param in create_pc, here we only fetch first one
	fac, err := h.faction.Fetch(ctx, faction.Filter{
		WorldID: ro.WorldID,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch world faction")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert entity backup
	bu := entity.Backup{
		ID:          entityID,
		UserID:      u.ID,
		CellID:      c.CellID,
		FactionID:   fac.ID,
		X:           0,
		Y:           0,
		Rot:         0,
		Radius:      eTemplate.Radius,
		At:          time.Now().UnixNano(),
		AnimationID: idle,
		AnimationAt: 0,
		Objects:     eTemplate.Objects,
		// Name:        "",
		// Stats: eTemplate.Stats,
		// Effects: eTemplate.Effects,
	}
	if err := h.entity.InsertBackup(ctx, bu); err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Insert PC
	pc := entity.PC{
		ID:       ulid.NewID(),
		EntityID: bu.ID,
		UserID:   u.ID,
		WorldID:  ro.WorldID,
	}
	if err := h.entity.InsertPC(ctx, pc); err != nil {
		logger.Error().Err(err).Msg("failed to create pc")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	if err := h.faction.InsertPC(ctx, faction.PC{
		ID:         pc.ID,
		FactionID:  fac.ID,
		Permission: int64(faction.Member),
	}); err != nil {
		logger.Error().Err(err).Msg("failed to create faction pc")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	abilities, err := h.entity.CreateDefaultAbilities(ctx, pc.EntityID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create default abilities")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Assign 5 first abilities (maximum) to first hotbar
	pcp := entity.PCPreferences{
		ID:             ulid.NewID(),
		PCID:           pc.ID,
		AbilityHotbars: make(map[string]ulid.ID),
	}
	for i := 0; i < len(abilities) && i < 6; i++ {
		pcp.AbilityHotbars[fmt.Sprintf("hotkey-0-%d", i)] = abilities[i].AbilityID
	}

	if err := h.entity.InsertPCPreferences(ctx, pcp); err != nil {
		logger.Error().Err(err).Msg("failed to insert pc preferences")

		return &entity.PC{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger.Info().Msg("success")

	return &pc, nil
}
