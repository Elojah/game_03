package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/geometry"
	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/tile"
	"github.com/elojah/game_03/pkg/tile/template"
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

	// Create island template and use it for spawns afterward
	wi := template.NewIslands(template.IslandsOptions{
		Width:         (cellWidth * width * 2) + 1,
		Height:        (cellHeight * height * 2) + 1,
		IslandWidth:   60,
		IslandHeight:  34,
		PaddingWidth:  20,
		PaddingHeight: 12,
		PathWidth:     15,
	})

	var g wang.Grid

	// Generate with wang constraints
	// Heuristic order is important
	g.Generate(ts.WangSets[0], cellHeight*height, cellWidth*width, wi.Heuristic(), template.AttunedHeuristic)

	collisions := tile.ObjectsByGID(ts.Tiles[0].ObjectGroup.Objects)

	// TODO: add goroutine pool
	for i, cl := range cells {
		for j, c := range cl {
			rect := geometry.Rect{
				Origin: geometry.Vec2{
					X: int64(j) * cellWidth,
					Y: int64(i) * cellHeight,
				},
				Width:  uint64(cellWidth),
				Height: uint64(cellHeight),
			}

			m, err := g.Tilemap(rect, ts, collisions)
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

			// assign cell id to waypoints
			wps := wi.Waypoints(rect, c.ID)

			if err := h.room.PopulateWaypoints(ctx, wps); err != nil {
				logger.Error().Err(err).Msg("failed to populate waypoints")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}

			logger.Info().Str("tilemap_id", c.Tilemap.String()).Str("cell_id", c.ID.String()).Msg("cell created")
		}
	}

	logger.Info().Msg("success")

	return &types.StringValue{Value: w.ID.String()}, nil
}
