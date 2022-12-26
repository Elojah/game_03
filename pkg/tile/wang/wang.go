package wang

import (
	"fmt"

	gtile "github.com/elojah/game_03/pkg/tile"
)

type edge uint8

const (
	top edge = iota
	right
	down
	left
)

type id int

type color byte

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

func generate(w gtile.WangSet, height int64, width int64) [][]id { //nolint: gocognit
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
			top:   colors(wt.WangID[7] + wt.WangID[0] + wt.WangID[1]),
			right: colors(wt.WangID[1] + wt.WangID[2] + wt.WangID[3]),
			down:  colors(wt.WangID[3] + wt.WangID[4] + wt.WangID[5]),
			left:  colors(wt.WangID[5] + wt.WangID[6] + wt.WangID[7]),
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

	g := make([][]id, height)
	for i := range g {
		g[i] = make([]id, width)
	}

	for i := int64(0); i < height; i++ {
		for j := int64(0); j < width; j++ {
			constraints := make(map[edge]colors)

			// set only top & left constraints for usual top left run
			if i > 0 {
				t, ok := ts[g[i-1][j]]

				if !ok {
					// should never happen, it means we set an invalid id tile in a previous loop
					fmt.Println("FAILED I with: ", i-1, j, g[i-1][j])
					continue
				}

				constraints[top] = t[down]
			}

			if j > 0 {
				t, ok := ts[g[i][j-1]]

				if !ok {
					// should never happen, it means we set an invalid id tile in a previous loop
					fmt.Println("FAILED J with: ", i, j-1, g[i][j-1])
					continue
				}

				constraints[left] = t[right]
			}

			if len(constraints) == 0 {
				// no constraints (first tile to put), put random one
				g[i][j] = ts.rand()

				continue
			}

			var candidates map[id]struct{}

			filter := false

			for o, constraint := range constraints {
				c, ok := ot[o][constraint]
				if !ok {
					// constraint color not found, return invalid index
					fmt.Println("single constraint not found")

					g[i][j] = 0

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
				g[i][j] = k

				break
			}
		}
	}

	return g
}
