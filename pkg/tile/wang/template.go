package wang

import (
	"math/rand"
	"time"

	"github.com/elojah/game_03/pkg/geometry"
	gtile "github.com/elojah/game_03/pkg/tile"
)

type Template [][]color

func NewTemplate(w gtile.WangSet, height int64, width int64) Template {
	rand.Seed(time.Now().UnixNano())

	t := make(Template, (2*height)+1) //nolint: gomnd

	for i := 0; i < len(t); i++ {
		t[i] = make([]color, (2*width)+1) //nolint: gomnd
		for j := 0; j < len(t[i]); j++ {
			if i == 0 || j == 0 ||
				i == len(t)-1 || j == len(t[i])-1 {
				t[i][j] = 0 // void
			} else if i == 1 || j == 1 ||
				i == len(t)-2 || j == len(t[i])-2 {
				t[i][j] = 3 // fence
			} else {
				t[i][j] = 0 // void
			}
		}
	}

	nPlatforms := int64(1000)
	for n := int64(0); n < nPlatforms; n++ {
		p := geometry.Vec2{
			X: rand.Int63n(width - 1),
			Y: rand.Int63n(height - 1),
		}
		w := 10 + rand.Int63n(100)
		h := 10 + rand.Int63n(100)

		for i := p.Y; i < p.Y+h && i < int64(len(t)); i++ {
			for j := p.X; j < p.X+w && j < int64(len(t[i])); j++ {
				if i == p.Y || j == p.X ||
					i == p.Y+h-1 || j == p.X+w-1 {
					t[i][j] = 0 // void
				} else if i == p.Y+1 || j == p.X+1 ||
					i == p.Y+h-2 || j == p.X+w-2 {
					t[i][j] = 3 // fence
				} else {
					t[i][j] = 1 // grass
				}
			}
		}
	}

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
