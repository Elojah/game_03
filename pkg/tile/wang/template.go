package wang

import gtile "github.com/elojah/game_03/pkg/tile"

type Template [][]color

func NewTemplate(w gtile.WangSet, height int64, width int64) Template {
	t := make(Template, (3*height)+1) //nolint: gomnd
	for i := 0; i < len(t); i++ {
		t[i] = make([]color, (3*width)+1) //nolint: gomnd
		for j := 0; j < len(t[i]); j++ {
			if i == 0 || j == 0 ||
				i == len(t)-1 || j == len(t[i])-1 {
				t[i][j] = 3 // fence

				continue
			}

			if i == j {
				t[i][j] = 3 // fence

				continue
			}

			t[i][j] = 1
		}
	}

	return t
}

func (t Template) Heuristic() Heuristic {
	return func(candidates map[id]struct{}, y int64, x int64, ts tiles) id {
		var result id
		var resultMatches int

		for c := range candidates {
			var matches int

			// top left
			if color(ts[c][top][0]) == t[(3 * y)][(3*x)] {
				matches++
			}

			// top
			if color(ts[c][top][1]) == t[(3 * y)][(3*x)+1] {
				matches++
			}

			// top right
			if color(ts[c][top][2]) == t[(3 * y)][(3*x)+2] {
				matches++
			}

			// left
			if color(ts[c][left][1]) == t[(3*y)+1][(3*x)] {
				matches++
			}

			// right
			if color(ts[c][right][1]) == t[(3*y)+1][(3*x)+2] {
				matches++
			}

			// down left
			if color(ts[c][left][2]) == t[(3*y)+2][(3*x)] {
				matches++
			}

			// down
			if color(ts[c][down][1]) == t[(3*y)+2][(3*x)+1] {
				matches++
			}

			// down right
			if color(ts[c][down][2]) == t[(3*y)+2][(3*x)+2] {
				matches++
			}

			if matches > resultMatches {
				resultMatches = matches
				result = c
			}
		}

		return result
	}
}
