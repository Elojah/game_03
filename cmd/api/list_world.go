package main

import (
	"context"
	"errors"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/entity/dto"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ListWorld(ctx context.Context, req *dto.ListPCReq) (*dto.ListPCResp, error) {
	logger := log.With().Str("method", "list_pc").Logger()

	if req == nil {
		return &dto.ListPCResp{}, status.New(codes.Internal, gerrors.ErrNullRequest{}.Error()).Err()
	}

	// #Authenticate
	ses, err := h.user.Auth(ctx)
	if err != nil {
		return &dto.ListPCResp{}, status.New(codes.Unauthenticated, err.Error()).Err()
	}

	// #Fetch rooms
	r, err := h.room.Fetch(ctx,
		room.Filter{
			ID: &req.RoomID,
		},
	)
	if err != nil {
		if errors.As(err, &gerrors.ErrNotFound{}) {
			logger.Error().Err(err).Msg("room not found")

			return &dto.ListPCResp{}, status.New(codes.NotFound, err.Error()).Err()
		}

		logger.Error().Err(err).Msg("failed to fetch rooms")

		return &dto.ListPCResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch pcs
	pcs, state, err := h.entity.FetchManyPC(ctx,
		entity.FilterPC{
			UserID:  &ses.UserID,
			WorldID: &r.WorldID,
			Size:    int(req.Size_),
			State:   req.State,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch pcs")

		return &dto.ListPCResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	// #Fetch entities
	entityIDs := make([]ulid.ID, 0, len(pcs))
	for _, pc := range pcs {
		entityIDs = append(entityIDs, pc.EntityID)
	}

	ents, _, err := h.entity.FetchManyBackup(ctx, entity.Filter{
		IDs:  entityIDs,
		Size: len(entityIDs),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch entities")

		return &dto.ListPCResp{}, status.New(codes.Internal, err.Error()).Err()
	}

	entsMap := make(map[ulid.ID]entity.E, len(ents))
	for _, ent := range ents {
		entsMap[ent.ID] = ent
	}

	// #Associate entities to PCs
	result := dto.ListPCResp{PCs: make([]dto.PC, 0, len(pcs)), State: state}

	for _, pc := range pcs {
		ent, ok := entsMap[pc.EntityID]
		if !ok {
			// should not happen
			continue
		}

		result.PCs = append(result.PCs, dto.PC{
			PC:     pc,
			Entity: ent,
		})
	}

	logger.Info().Msg("success")

	return &result, nil
}
