package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/elojah/game_03/pkg/room"
	"github.com/elojah/game_03/pkg/tile"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateWorld creates a new randomly generated world setup for a new room.
func (h *handler) CreateWorld(ctx context.Context, req *types.Empty) (*types.StringValue, error) {
	logger := log.With().Str("method", "create_world").Logger()

	params := tile.Params{
		Height:          32,    //nolint: gomnd
		Width:           32,    //nolint: gomnd
		CellHeight:      34,    //nolint: gomnd
		CellWidth:       60,    //nolint: gomnd
		WorldDensity:    0.015, //nolint: gomnd
		PlatformDensity: 0.99,  //nolint: gomnd
		PlatformPadding: 0.99,  //nolint: gomnd
		SizeMin:         5,     //nolint: gomnd
		SizeMax:         15,    //nolint: gomnd
		Distortion:      0,
		PathMin:         0,
		PathMax:         0,
		PathDistorsion:  0,

		Set: tile.Set{
			Columns:     2, //nolint: gomnd
			FirstGID:    1,
			Image:       "../img/01GDB3DDM9Q1R1XTSSNNB9CYJV.png",
			ImageHeight: 32, //nolint: gomnd
			ImageWidth:  64, //nolint: gomnd
			Margin:      0,
			Name:        "01GDB3DDM9Q1R1XTSSNNB9CYJV",
			Spacing:     0,
			TileCount:   2,  //nolint: gomnd
			TileHeight:  32, //nolint: gomnd
			TileWidth:   32, //nolint: gomnd
		},
	}

	// #Create world
	w := room.World{
		ID:         ulid.NewID(),
		Height:     int64(params.Height),
		Width:      int64(params.Width),
		CellHeight: int64(params.CellHeight) * 32, // TODO: figure out naming here cause it means different than above (x tile_size)
		CellWidth:  int64(params.CellWidth) * 32,  // TODO: figure out naming here cause it means different than above (x tile_size)
	}
	if err := h.room.InsertWorld(ctx, w); err != nil {
		logger.Error().Err(err).Msg("failed to create world")

		return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	}

	logger = logger.With().Str("world_id", w.ID.String()).Logger()

	// #Create world cells
	cells := w.NewCells()

	g := tile.GroundGenerator{}
	g.Gen(params)

	// TODO: add goroutine pool
	for i, cl := range cells {
		for j, c := range cl {
			m := g.Tilemap(params, uint64(i), uint64(j))

			raw, err := json.Marshal(m)
			if err != nil {
				logger.Error().Err(err).Msg("failed to marshal map")

				return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			}

			id := ulid.NewID()

			f, err := os.Create("cmd/browser/dist/json/" + id.String() + ".json")
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
