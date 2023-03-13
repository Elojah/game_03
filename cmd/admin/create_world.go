package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/geometry"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/tile"
	"github.com/elojah/game_03/pkg/tile/wang"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateWorld creates a new randomly generated world setup for a new room.
func (h *handler) CreateWorld(ctx context.Context, req *types.Empty) (*types.StringValue, error) {
	logger := log.With().Str("method", "create_world").Logger()

	height := int64(3)      //nolint: gomnd
	width := int64(3)       //nolint: gomnd
	cellHeight := int64(34) //nolint: gomnd
	cellWidth := int64(60)  //nolint: gomnd

	id := ulid.NewID()

	// #Create world
	w := room.World{
		ID:         id,
		Name:       "template_" + id.String(),
		Height:     height,
		Width:      width,
		CellHeight: cellHeight * 32, // TODO: figure out naming here cause it means different than above (* tile_size)
		CellWidth:  cellWidth * 32,  // TODO: figure out naming here cause it means different than above (* tile_size)
	}
	if err := h.room.InsertWorld(ctx, w); err != nil {
		logger.Error().Err(err).Msg("failed to create world")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger = logger.With().Str("world_id", w.ID.String()).Logger()

	// #Create world cells
	cells := w.NewCells()

	// TODO: tile set in parameter ?
	tsID := ulid.MustParse("01GMAP5JY8YRHZJ45TRWZA8VHM")

	ts, err := h.tile.FetchSet(ctx, tile.FilterSet{ID: tsID})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch tileset")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	if len(ts.WangSets) == 0 {
		err := errors.ErrMissingWangSet{ID: tsID.String()}
		logger.Error().Err(err).Msg("missing wangset in tileset")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	if len(ts.Tiles) == 0 {
		err := errors.ErrMissingCollisionLayer{ID: tsID.String()}
		logger.Error().Err(err).Msg("missing collision layer in tileset")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	// Create island template
	wi := wang.NewIslands(ts.WangSets[0], cellHeight*height, cellWidth*width)

	var g wang.Grid

	// Generate with wang constraints
	g.Generate(ts.WangSets[0], cellHeight*height, cellWidth*width, wi.Heuristic())
	// g.GenerateFlat(ts.WangSets[0], cellHeight*height, cellWidth*width)

	collisions := tile.ObjectsByGID(ts.Tiles[0].ObjectGroup.Objects)

	// Fetch and create altar animation
	// ________________________________

	altar, err := h.entity.FetchTemplate(ctx, entity.FilterTemplate{
		// Pick first result (random) of altar template
		// TODO: Add as parameter to pick specific Altar
		Name: func(s string) *string { return &s }("Altar"),
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch altar template")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	ans, _, err := h.entity.FetchManyAnimation(ctx, entity.FilterAnimation{
		EntityID: altar.ID,
		Size:     1,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to fetch many animation")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	if len(ans) == 0 {
		err := errors.ErrMissingDefaultAnimations{EntityID: altar.ID.String()}

		logger.Error().Err(err).Msg("missing animation")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	entityID := ulid.NewID()

	var main ulid.ID

	for _, an := range ans {
		an.ID = ulid.NewID()
		an.EntityID = entityID

		// default animation name is always "main"
		if an.Name == "main" {
			main = an.ID
		}

		if err := h.entity.InsertAnimation(ctx, an); err != nil {
			logger.Error().Err(err).Msg("failed to create animation")

			return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
		}
	}

	// #Insert entity
	e := entity.E{
		ID:          entityID,
		UserID:      ulid.NewID(),
		CellID:      cells[0][0].ID,
		X:           0,
		Y:           0,
		Rot:         0,
		Radius:      10, //nolint: gomnd
		At:          time.Now().UnixNano(),
		AnimationID: main,
		AnimationAt: 0,
	}
	if err := h.entity.Insert(ctx, e); err != nil {
		logger.Error().Err(err).Msg("failed to create entity")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	// ________________________________
	// !Fetch and create altar animation

	// TODO: add goroutine pool
	for i, cl := range cells {
		for j, c := range cl {
			m, err := g.Tilemap(geometry.Rect{
				Origin: geometry.Vec2{
					X: int64(j) * cellWidth,
					Y: int64(i) * cellHeight,
				},
				Width:  uint64(cellWidth),
				Height: uint64(cellHeight),
			}, ts, collisions)
			if err != nil {
				logger.Error().Err(err).Msg("failed to create tilemap")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}

			raw, err := json.Marshal(m)
			if err != nil {
				logger.Error().Err(err).Msg("failed to marshal map")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}

			id := ulid.NewID()

			f, err := os.Create("cmd/client/dist/json/assets/" + id.String() + ".json")
			if err != nil {
				logger.Error().Err(err).Msg("failed to create file")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}

			if _, err := f.Write(raw); err != nil {
				logger.Error().Err(err).Msg("failed to write file")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}

			c.Tilemap = id
			c.X = int64(j)
			c.Y = int64(i)

			if err := h.room.InsertCell(ctx, c); err != nil {
				logger.Error().Err(err).Msg("failed to create cell")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}

			if err := h.room.InsertWorldCell(ctx, room.WorldCell{
				WorldID: w.ID,
				CellID:  c.ID,
				X:       c.X,
				Y:       c.Y,
			}); err != nil {
				logger.Error().Err(err).Msg("failed to create world cell")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}

			logger.Info().Str("tilemap_id", c.Tilemap.String()).Str("cell_id", c.ID.String()).Msg("cell created")
		}
	}

	logger.Info().Msg("success")

	return &types.StringValue{Value: w.ID.String()}, nil
}
