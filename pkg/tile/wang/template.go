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
	fence = 3

	templatePathWidth = 10
)

type Template [][]color

func NewTemplate(w gtile.WangSet, height int64, width int64) Template { //nolint: gocognit
	rand.Seed(time.Now().UnixNano())

	t := make(Template, (2*height)+1) //nolint: gomnd

	for i := 0; i < len(t); i++ {
		t[i] = make([]color, (2*width)+1) //nolint: gomnd
	}

	randPlatforms := width * height / 100
	nPlatforms := randPlatforms + rand.Int63n(randPlatforms)
	for n := int64(0); n < nPlatforms; n++ {
		p := geometry.Vec2{
			X: rand.Int63n(width - 1),
			Y: rand.Int63n(height - 1),
		}
		w := ((randPlatforms) + rand.Int63n(randPlatforms)) / 10
		h := ((randPlatforms) + rand.Int63n(randPlatforms)) / 10

		for i := p.Y; i < p.Y+h && i < int64(len(t)); i++ {
			for j := p.X; j < p.X+w && j < int64(len(t[i])); j++ {

				// basic platform template
				if i == p.Y || j == p.X ||
					i == p.Y+h-1 || j == p.X+w-1 {
					t[i][j] = void
				} else if i == p.Y+1 || j == p.X+1 ||
					i == p.Y+h-2 || j == p.X+w-2 {
					t[i][j] = fence
				} else {
					t[i][j] = grass
				}

				// ensure path width are >= templatePathWidth
				if t[i][j] == grass {
					continue
				}

				if i > 1 && t[i-1][j] == grass {
					for ii := int64(0); i-ii > 0; i++ {
						if ii >= templatePathWidth {
							break
						}
						if t[i-ii][j] != grass {
							t[i][j] = grass
							break
						}
					}
				}

				if t[i][j] == grass {
					continue
				}

				if j > 1 && t[i][j-1] == grass {
					for jj := int64(0); j-jj > 0; j++ {
						if jj >= templatePathWidth {
							break
						}
						if t[i][j-jj] != grass {
							t[i][j] = grass
							break
						}
					}
				}

				if t[i][j] == grass {
					continue
				}

				if i < p.Y-1 && t[i+1][j] == grass {
					for ii := int64(0); i+ii < p.Y-1; i++ {
						if ii >= templatePathWidth {
							break
						}
						if t[i+ii][j] != grass {
							t[i][j] = grass
							break
						}
					}
				}

				if t[i][j] == grass {
					continue
				}

				if j < p.X-1 && t[i][j+1] == grass {
					for jj := int64(0); j+jj > 0; j++ {
						if jj >= templatePathWidth {
							break
						}
						if t[i][j+jj] != grass {
							t[i][j] = grass
							break
						}
					}
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
