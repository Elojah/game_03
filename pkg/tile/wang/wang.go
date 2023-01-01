package wang

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/geometry"
	gtile "github.com/elojah/game_03/pkg/tile"
)

// const (
// 	Corner = "corner"
// 	Edge   = "edge"
// 	All    = "all"
// )

type edge uint8

const (
	top edge = iota
	right
	down
	left
)

type id int

type color byte

// string method for debug purpose.
func (cs colors) string() string {
	b := strings.Builder{}
	b.WriteByte('_')

	for _, c := range cs {
		b.WriteString(strconv.Itoa(int(c)))
	}

	b.WriteByte('_')

	return b.String()
}

type colors string

type tile map[edge]colors

type tiles map[id]tile

func (ts tiles) rand() id {
	for id := range ts {
		// return first iteration, using random map access iteration from go specification
		// to fetch a random element
		return id
	}

	return 0
}

type orientedTile map[colors]map[id]struct{}

type orientedTiles map[edge]orientedTile

type Grid [][]id

func (g Grid) Tilemap(r geometry.Rect, ts gtile.Set) (gtile.Map, error) {
	if r.Origin.Y < 0 || r.Origin.Y+int64(r.Height) > int64(len(g)) ||
		r.Origin.X < 0 || r.Origin.X+int64(r.Width) > int64(len(g[r.Origin.Y])) {
		return gtile.Map{}, errors.ErrInvalidDimension{}
	}

	m := gtile.NewMap()

	m.Height = int(r.Height)
	m.Width = int(r.Width)
	m.NextLayerID = 3
	m.NextObjectID = 1
	m.Tilesets = append(m.Tilesets, ts.CleanCopy())

	// main layer
	layer := gtile.NewLayer()
	layer.ID = 1
	layer.Height = m.Height
	layer.Width = m.Width

	// collision layer
	clayer := gtile.NewLayer()
	clayer.ID = 2
	clayer.Height = m.Height
	clayer.Width = m.Width

	size := layer.Height * layer.Width
	data := make([]byte, 0, 4*size) //nolint: gomnd

	for i := r.Origin.Y; i < r.Origin.Y+int64(r.Height); i++ {
		for j := r.Origin.X; j < r.Origin.X+int64(r.Width); j++ {
			// collideData = binary.LittleEndian.AppendUint32(collideData, func(f Field) wangid {
			// 	if f.Collides() {
			// 		return 1
			// 	}
			// 	return 0
			// }(f))
			data = binary.LittleEndian.AppendUint32(data, uint32(g[i][j]))
		}
	}

	layer.Data = base64.StdEncoding.EncodeToString(data)

	m.Layers = append(m.Layers, layer)

	return m, nil
}

func (g *Grid) Generate(w gtile.WangSet, height int64, width int64) { //nolint: gocognit
	ts := make(tiles, len(w.WangTiles))
	ot := orientedTiles{
		top:   make(orientedTile),
		right: make(orientedTile),
		down:  make(orientedTile),
		left:  make(orientedTile),
	}

	for _, wt := range w.WangTiles {
		// discard wang tiles with missing information
		if len(wt.WangID) < 8 { //nolint: gomnd
			continue
		}

		t := tile{
			top:   colors([]byte{wt.WangID[7], wt.WangID[0], wt.WangID[1]}), // left to right
			right: colors([]byte{wt.WangID[1], wt.WangID[2], wt.WangID[3]}), // top to down
			down:  colors([]byte{wt.WangID[5], wt.WangID[4], wt.WangID[3]}), // left to right
			left:  colors([]byte{wt.WangID[7], wt.WangID[6], wt.WangID[5]}), // top to down
		}

		// add tiles to general usage tiles struct
		ts[id(wt.TileID)] = t

		// add tiles to oriented struct for fast access during populate
		if _, ok := ot[top][t[top]]; !ok {
			ot[top][t[top]] = map[id]struct{}{id(wt.TileID): {}}
		} else {
			ot[top][t[top]][id(wt.TileID)] = struct{}{}
		}

		if _, ok := ot[right][t[right]]; !ok {
			ot[right][t[right]] = map[id]struct{}{id(wt.TileID): {}}
		} else {
			ot[top][t[top]][id(wt.TileID)] = struct{}{}
		}

		if _, ok := ot[down][t[down]]; !ok {
			ot[down][t[down]] = map[id]struct{}{id(wt.TileID): {}}
		} else {
			ot[down][t[down]][id(wt.TileID)] = struct{}{}
		}

		if _, ok := ot[left][t[left]]; !ok {
			ot[left][t[left]] = map[id]struct{}{id(wt.TileID): {}}
		} else {
			ot[left][t[left]][id(wt.TileID)] = struct{}{}
		}
	}

	result := make(Grid, height)
	for i := range result {
		result[i] = make([]id, width)
	}

	for i := int64(0); i < height; i++ {
		for j := int64(0); j < width; j++ {
			constraints := make(map[edge]colors)

			// set only top & left constraints for usual top left run
			if i > 0 {
				t, ok := ts[result[i-1][j]]

				if !ok {
					// should never happen, it means we set an invalid id tile in a previous loop
					continue
				}

				constraints[top] = t[down]
			}

			if j > 0 {
				t, ok := ts[result[i][j-1]]

				if !ok {
					// should never happen, it means we set an invalid id tile in a previous loop
					continue
				}

				constraints[left] = t[right]
			}

			if len(constraints) == 0 {
				// no constraints (first tile to put), put random one
				result[i][j] = ts.rand()

				continue
			}

			var candidates map[id]struct{}

			filter := false

			for o, constraint := range constraints {
				c, ok := ot[o][constraint]
				if !ok {
					// constraint color not found, return invalid index
					fmt.Println("single constraint not found")

					result[i][j] = 0

					break
				}

				if !filter {
					// first pass, add all ids candidates
					candidates = c
					filter = true
				} else {
					// second pass, if id also exists in candidates, add as solution
					tmp := make(map[id]struct{})

					for k := range c {
						if _, ok := candidates[k]; ok {
							tmp[k] = struct{}{}
						}
					}

					candidates = tmp
				}
			}

			if len(candidates) == 0 {
				fmt.Println("constraints not found")
			}

			for k := range candidates {
				// set first in loop using random map access from golang specification
				result[i][j] = k

				break
			}
		}
	}

	*g = result
}
