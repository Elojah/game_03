package tile

import (
	"sort"

	"github.com/elojah/game_03/pkg/geometry"
)

type Objects []Object

// splitGrid split objects into integer grid.
func (os Objects) splitGrid() ([]float64, []float64, [][]bool) {
	var xSplit, ySplit []float64

	xUnique := make(map[float64]struct{})
	yUnique := make(map[float64]struct{})

	// remove duplicates
	for _, o := range os {
		if _, ok := xUnique[o.X]; !ok {
			xSplit = append(xSplit, o.X)
			xUnique[o.X] = struct{}{}
		}

		if _, ok := xUnique[o.X+o.Width]; !ok {
			xSplit = append(xSplit, o.X+o.Width)
			xUnique[o.X+o.Width] = struct{}{}
		}

		if _, ok := yUnique[o.Y]; !ok {
			ySplit = append(ySplit, o.Y)
			yUnique[o.Y] = struct{}{}
		}

		if _, ok := yUnique[o.Y+o.Height]; !ok {
			ySplit = append(ySplit, o.Y+o.Height)
			yUnique[o.Y+o.Height] = struct{}{}
		}
	}

	sort.Float64s(xSplit)
	sort.Float64s(ySplit)

	if len(xSplit) < 2 || len(ySplit) < 2 {
		return nil, nil, nil
	}

	values := make([][]bool, len(ySplit)-1)
	for i := range values {
		values[i] = make([]bool, len(xSplit)-1)
	}

	for _, o := range os {
		// find x range iterations in grid
		var xMin, xMax int64

		for i, x := range xSplit {
			if x == o.X {
				xMin = int64(i)
			}

			if x == o.X+o.Width {
				xMax = int64(i)

				break
			}
		}

		// find y range iterations in grid
		var yMin, yMax int64

		for i, y := range ySplit {
			if y == o.Y {
				yMin = int64(i)
			}

			if y == o.Y+o.Height {
				yMax = int64(i)

				break
			}
		}

		for i := yMin; i < yMax; i++ {
			for j := xMin; j < xMax; j++ {
				values[i][j] = true
			}
		}
	}

	return xSplit, ySplit, values
}

// dissectRectangles implements naive algorithm and not optimal solution
// to find minimum dissection of rectangle.
func dissectRectangles(grid [][]bool) []geometry.Rect { //nolint: gocognit
	var result []geometry.Rect

	findNext := func(grid [][]bool) (geometry.Rect, bool) {
		for i, line := range grid {
			for j, val := range line {
				if val {
					r := geometry.Rect{
						Origin: geometry.Vec2{
							X: int64(j),
							Y: int64(i),
						},
						Width:  1,
						Height: 1,
					}

					for jj := j + 1; jj < len(grid[i]); jj++ {
						if !grid[i][jj] {
							break
						}
						r.Width++
					}

					var incomplete bool

					for ii := i + 1; ii < len(grid); ii++ {
						for jj := j; uint64(jj) < uint64(j)+r.Width; jj++ {
							if !grid[ii][jj] {
								incomplete = true

								break
							}
						}

						if incomplete {
							break
						}
						r.Height++
					}

					return r, true
				}
			}
		}

		return geometry.Rect{}, false
	}

	for r, ok := findNext(grid); ok; r, ok = findNext(grid) {
		result = append(result, r)

		// add rectangle to grid
		for i := r.Origin.Y; i < r.Origin.Y+int64(r.Height); i++ {
			for j := r.Origin.X; j < r.Origin.X+int64(r.Width); j++ {
				grid[i][j] = false
			}
		}
	}

	return result
}

func (os Objects) MinimumDissection() Objects {
	xSplit, ySplit, values := os.splitGrid()

	rs := dissectRectangles(values)

	result := make([]Object, 0, len(rs))
	for _, r := range rs {
		result = append(result, Object{
			X:      xSplit[r.Origin.X],
			Y:      ySplit[r.Origin.Y],
			Width:  xSplit[r.Origin.X+int64(r.Width)] - xSplit[r.Origin.X],
			Height: ySplit[r.Origin.Y+int64(r.Height)] - ySplit[r.Origin.Y],
		})
	}

	return result
}
