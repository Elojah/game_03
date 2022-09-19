package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/elojah/game_03/pkg/tile"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/gogo/protobuf/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TMP TEST
func (h *handler) CreateTilemap(ctx context.Context, req *types.Empty) (*types.StringValue, error) {
	logger := log.With().Str("method", "create_tilemap").Logger()

	g := tile.GroundGenerator{}
	// params := tile.Params{
	// 	Height:          3,
	// 	Width:           3,
	// 	CellHeight:      11,
	// 	CellWidth:       20,
	// 	WorldDensity:    0.9,
	// 	PlatformDensity: 0.7,
	// 	PlatformPadding: 0.9,
	// 	SizeMin:         5,
	// 	SizeMax:         15,
	// 	Distortion:      0,
	// 	PathMin:         0,
	// 	PathMax:         0,
	// 	PathDistorsion:  0,

	// 	Set: tile.Set{
	// 		Columns:     64,
	// 		FirstGID:    1,
	// 		Image:       "../img/01FG1V2RERMABQRXEYKT8T1HYV.png",
	// 		ImageHeight: 1024,
	// 		ImageWidth:  1024,
	// 		Margin:      0,
	// 		Name:        "01FG1V2RERMABQRXEYKT8T1HYV",
	// 		Spacing:     0,
	// 		TileCount:   4096,
	// 		TileHeight:  16,
	// 		TileWidth:   16,
	// 	},
	// }

	params := tile.Params{
		Height:          1,
		Width:           1,
		CellHeight:      200,
		CellWidth:       200,
		WorldDensity:    0.9,
		PlatformDensity: 0.7,
		PlatformPadding: 0.9,
		SizeMin:         5,
		SizeMax:         15,
		Distortion:      0,
		PathMin:         0,
		PathMax:         0,
		PathDistorsion:  0,

		Set: tile.Set{
			Columns:     1,
			FirstGID:    1,
			Image:       "../img/01GDB3DDM9Q1R1XTSSNNB9CYJV.png",
			ImageHeight: 32,
			ImageWidth:  32,
			Margin:      0,
			Name:        "01GDB3DDM9Q1R1XTSSNNB9CYJV",
			Spacing:     0,
			TileCount:   1,
			TileHeight:  1,
			TileWidth:   1,
		},
	}

	g.Gen(params)
	m := g.Tilemap(params)

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

	logger.Info().Msg("success")

	return &types.StringValue{Value: id.String()}, nil
}
