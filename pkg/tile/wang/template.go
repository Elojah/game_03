package wang

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/elojah/game_03/pkg/geometry"
	gtile "github.com/elojah/game_03/pkg/tile"
)

const (
	void  = 0
	grass = 1
	wall  = 2
	fence = 3

	templatePathWidth = 3

	// debug variable, remove me
	tmpAdjust = 1
)

type Template [][]color

func NewTemplate(w gtile.WangSet, height int64, width int64) Template { //nolint: gocognit
	rand.Seed(time.Now().UnixNano())

	width = (2 * width) + 1
	height = (2 * height) + 1

	t := make(Template, height) //nolint: gomnd

	for i := 0; i < len(t); i++ {
		t[i] = make([]color, width) //nolint: gomnd
	}

	randPlatforms := width * height / (100 * tmpAdjust)
	nPlatforms := randPlatforms + rand.Int63n(randPlatforms)
	for n := int64(0); n < nPlatforms; n++ {
		p := geometry.Vec2{
			X: rand.Int63n(width - 1),
			Y: rand.Int63n(height - 1),
		}
		w := ((randPlatforms) + rand.Int63n(randPlatforms)) / (50 * tmpAdjust)
		h := ((randPlatforms) + rand.Int63n(randPlatforms)) / (50 * tmpAdjust)

		for i := p.Y; i < p.Y+h && i < int64(len(t)); i++ {
			for j := p.X; j < p.X+w && j < int64(len(t[i])); j++ {

				if t[i][j] != void {
					continue
				}

				// basic platform template
				if i == p.Y || j == p.X ||
					i == p.Y+h-1 || j == p.X+w-1 {
					t[i][j] = wall
				} else if i == p.Y+1 || j == p.X+1 ||
					i == p.Y+h-2 || j == p.X+w-2 {
					t[i][j] = fence
				} else {
					t[i][j] = grass
				}
			}
		}
	}

	// Set void around island.
	for i := int64(0); i < int64(len(t)); i++ {
		for j := int64(0); j < int64(len(t[i])); j++ {
			if t[i][j] != wall {
				break
			}

			t[i][j] = void
		}
	}

	for i := int64(0); i < int64(len(t)); i++ {
		for j := int64(len(t[i]) - 1); j > int64(0); j-- {
			if t[i][j] != wall {
				break
			}

			t[i][j] = void
		}
	}

	if len(t) > 0 {
		for j := int64(0); j < int64(len(t[0])); j++ {
			for i := int64(0); i < int64(len(t)); i++ {
				if t[i][j] != wall {
					break
				}

				t[i][j] = void
			}
		}
	}

	// 		// ensure path width are >= templatePathWidth
	// 		if t[i][j] == void {
	// 			continue
	// 		}

	// 		// horizontal from left to right
	// 		if j > 1 && t[i][j-1] == void {
	// 			for n := int64(0); j+n < int64(len(t[i])) && n <= templatePathWidth; n++ {
	// 				if t[i][j+n] == void {
	// 					t[i][j+n] = grass
	// 				}
	// 			}
	// 		}

	// 		// vertical from top to down
	// 		if i > 1 && t[i-1][j] == void {
	// 			for n := int64(0); i+n < int64(len(t)) && n <= templatePathWidth; n++ {
	// 				if t[i+n][j] == void {
	// 					t[i+n][j] = grass
	// 				}
	// 			}
	// 		}

	// 		// diagonal from top-left to down-right
	// 		if i > 1 && j > 1 && t[i-1][j-1] == void {
	// 			for n := int64(0); i+n < int64(len(t)) && j+n < int64(len(t[i+n])) && n <= templatePathWidth; n++ {
	// 				if n == 0 || n == templatePathWidth {
	// 					t[i+n][j+n] = fence
	// 				} else {
	// 					t[i+n][j+n] = grass
	// 				}
	// 			}
	// 		}

	// 		// diagonal from down-left to top-right
	// 		if i < int64(len(t)-1) && j > 1 && t[i+1][j-1] == void {
	// 			for n := int64(0); i-n > 0 && j+n < int64(len(t[i])) && n <= templatePathWidth; n++ {
	// 				if n == 0 || n == templatePathWidth {
	// 					t[i-n][j+n] = fence
	// 				} else {
	// 					t[i-n][j+n] = grass
	// 				}
	// 			}
	// 		}

	return t
}

func (t Template) Heuristic() Heuristic {
	return func(candidates map[id]struct{}, y int64, x int64, ts tiles) id {
		result := id(-1)
		resultMatches := 0

		for c := range candidates {
			var matches int

			// top left
			if color(ts[c][top][0]) == t[(2 * y)][(2*x)] {
				matches++
			}

			// top
			if color(ts[c][top][1]) == t[(2 * y)][(2*x)+1] {
				matches++
			}

			// top right
			if color(ts[c][top][2]) == t[(2 * y)][(2*x)+2] {
				matches++
			}

			// left
			if color(ts[c][left][1]) == t[(2*y)+1][(2*x)] {
				matches++
			}

			// right
			if color(ts[c][right][1]) == t[(2*y)+1][(2*x)+2] {
				matches++
			}

			// down left
			if color(ts[c][left][2]) == t[(2*y)+2][(2*x)] {
				matches++
			}

			// down
			if color(ts[c][down][1]) == t[(2*y)+2][(2*x)+1] {
				matches++
			}

			// down right
			if color(ts[c][down][2]) == t[(2*y)+2][(2*x)+2] {
				matches++
			}

			if matches > resultMatches || result == -1 {
				resultMatches = matches
				result = c
			}
		}

		return result
	}
}

func (t Template) String() string {
	var b strings.Builder

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			b.WriteString(strconv.Itoa(int(t[i][j])))
		}
		b.WriteString("\n")
	}

	return b.String()
}
