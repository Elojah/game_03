package template

import (
	"strconv"
	"strings"

	"github.com/elojah/game_03/pkg/geometry"
	"github.com/elojah/game_03/pkg/tile/wang"
)

const (
	void wang.Color = iota
	grass
	wall
	fence
)

type Islands [][]wang.Color

type Spawns []geometry.Vec2

type IslandsOptions struct {
	Width         int64
	Height        int64
	IslandWidth   int64
	IslandHeight  int64
	PaddingWidth  int64
	PaddingHeight int64
	PathWidth     int64
}

func NewIslands(opts IslandsOptions) Islands { //nolint: gocognit
	t := make(Islands, opts.Height) //nolint: gomnd

	for i := 0; i < len(t); i++ {
		t[i] = make([]wang.Color, opts.Width) //nolint: gomnd
	}

	cellHeight := opts.IslandHeight + opts.PaddingHeight
	cellWidth := opts.IslandWidth + opts.PaddingWidth
	nHeight := opts.Height / cellHeight
	nWidth := opts.Width / cellWidth

	for i := int64(0); i < nHeight; i++ {
		for j := int64(0); j < nWidth; j++ {
			// platform
			for ii := int64(0); ii < opts.IslandHeight; ii++ {
				for jj := int64(0); jj < opts.IslandWidth; jj++ {
					c := grass

					if ii == 0 || ii == opts.IslandHeight-1 ||
						jj == 0 || jj == opts.IslandWidth-1 {
						c = fence
					}

					t[(i*cellHeight)+ii+(opts.PaddingHeight/2)][(j*cellWidth)+jj+(opts.PaddingWidth/2)] = c
				}
			}

			// horizontal path
			for ii := int64(0); ii < opts.PathWidth; ii++ {
				for jj := int64(0); jj < cellWidth; jj++ {
					c := grass

					if ii == 0 || ii == opts.PathWidth-1 ||
						(jj == 0 && j == 0) || (jj == cellWidth-1 && j == nWidth-1) {
						c = fence
					}

					t[(i*cellHeight)+ii+((cellHeight/2)-(opts.PathWidth/2))][(j*cellWidth)+jj] = c
				}
			}

			// vertical path
			for ii := int64(0); ii < cellHeight; ii++ {
				for jj := int64(0); jj < opts.PathWidth; jj++ {
					c := grass

					if jj == 0 || jj == opts.IslandWidth-1 ||
						(ii == 0 && i == 0) || (ii == cellWidth-1 && i == nWidth-1) {

						c = fence
					}

					t[(i*cellHeight)+ii][(j*cellWidth)+jj+((cellWidth/2)-(opts.PathWidth/2))] = c
				}
			}

		}
	}

	return t
}

func (t Islands) Heuristic() wang.Heuristic {
	return func(candidates map[wang.ID]struct{}, y int64, x int64, ts wang.Tiles) map[wang.ID]struct{} {
		result := make(map[wang.ID]struct{})
		resultMatches := 0

		for c := range candidates {
			var matches int

			// top left
			if ts[c][wang.Top][0] == byte(t[(2 * y)][(2*x)]) {
				matches++
			}

			// top
			if ts[c][wang.Top][1] == byte(t[(2 * y)][(2*x)+1]) {
				matches++
			}

			// top right
			if ts[c][wang.Top][2] == byte(t[(2 * y)][(2*x)+2]) {
				matches++
			}

			// left
			if ts[c][wang.Left][1] == byte(t[(2*y)+1][(2*x)]) {
				matches++
			}

			// right
			if ts[c][wang.Right][1] == byte(t[(2*y)+1][(2*x)+2]) {
				matches++
			}

			// down left
			if ts[c][wang.Left][2] == byte(t[(2*y)+2][(2*x)]) {
				matches++
			}

			// down
			if ts[c][wang.Down][1] == byte(t[(2*y)+2][(2*x)+1]) {
				matches++
			}

			// down right
			if ts[c][wang.Down][2] == byte(t[(2*y)+2][(2*x)+2]) {
				matches++
			}

			if matches > resultMatches || len(result) == 0 {
				resultMatches = matches
				// reset map and assign single result
				result = map[wang.ID]struct{}{c: {}}
			} else if matches == resultMatches {
				// add another solution
				result[c] = struct{}{}
			}
		}

		return result
	}
}

func (t Islands) String() string {
	var b strings.Builder

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			b.WriteString(strconv.Itoa(int(t[i][j])))
		}
		b.WriteString("\n")
	}

	return b.String()
}
