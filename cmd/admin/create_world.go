package main

import (
	"context"
	"encoding/json"
	"os"

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

	height := int64(32)     //nolint: gomnd
	width := int64(32)      //nolint: gomnd
	cellHeight := int64(34) //nolint: gomnd
	cellWidth := int64(60)  //nolint: gomnd

	id := ulid.NewID()

	// #Create world
	w := room.World{
		ID:         id,
		Name:       "template_" + id.String(),
		Height:     height,
		Width:      width,
		CellHeight: cellHeight * 32, // TODO: figure out naming here cause it means different than above (x tile_size)
		CellWidth:  cellWidth * 32,  // TODO: figure out naming here cause it means different than above (x tile_size)
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

	var g wang.Grid

	g.Generate(ts.WangSets[0], cellHeight*height, cellWidth*width)
	// g.GenerateFlat(ts.WangSets[0], cellHeight*height, cellWidth*width)

	collisions := tile.ObjectsByGID(ts.Tiles[0].ObjectGroup.Objects)

	// an, err := h.entity.FetchAnimation(ctx, entity.FilterAnimation{
	// 	ID: ulid.MustParse("01FM1YY76G9879YJSZAFSGJ52Y"),
	// })
	// if err != nil {
	// 	logger.Error().Err(err).Msg("failed to fetch animation")

	// 	return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
	// }

	// now := time.Now()

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

			// decode world data (ok to instant test parsing) to find first non collides spot for spawn
			// data, err := base64.StdEncoding.DecodeString(m.Layers[1].Data)
			// for i := 0; i < len(data); i += 4 {
			// 	n := binary.LittleEndian.Uint32(data[i : i+4])
			// 	if !tile.Field(n).Collides() {
			// 		// a := an
			// 		// a.ID = ulid.NewID()
			// 		// a.EntityID = ulid.NewID()

			// 		// if err := h.entity.InsertAnimation(ctx, a); err != nil {
			// 		// 	logger.Error().Err(err).Msg("failed to create animation")

			// 		// 	return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			// 		// }

			// 		// e := entity.E{
			// 		// 	ID:          a.EntityID,
			// 		// 	UserID:      nil,
			// 		// 	CellID:      c.ID,
			// 		// 	Name:        fmt.Sprintf("spawn_%d_%d", i, j),
			// 		// 	X:           int64(i) % w.CellWidth,
			// 		// 	Y:           int64(i) / w.CellWidth,
			// 		// 	Rot:         0,
			// 		// 	Radius:      0,
			// 		// 	At:          now.Unix(),
			// 		// 	AnimationID: a.ID,
			// 		// 	AnimationAt: 0,
			// 		// }

			// 		// if err := h.entity.InsertNPC(ctx, entity.NPC{
			// 		// 	ID:       ulid.NewID(),
			// 		// 	WorldID:  w.ID,
			// 		// 	EntityID: e.ID,
			// 		// }); err != nil {
			// 		// 	logger.Error().Err(err).Msg("failed to create npc")

			// 		// 	return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			// 		// }

			// 		// if err := h.entity.InsertBackup(ctx, e); err != nil {
			// 		// 	logger.Error().Err(err).Msg("failed to create entity backup")

			// 		// 	return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			// 		// }

			// 		// if err := h.entity.Insert(ctx, e); err != nil {
			// 		// 	logger.Error().Err(err).Msg("failed to create entity")

			// 		// 	return &types.StringValue{}, status.New(codes.Internal, err.Error()).Err()
			// 		// }

			// 		break
			// 	}
			// }

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
