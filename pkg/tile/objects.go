package tile

import (
	"sort"

	"github.com/elojah/game_03/pkg/geometry"
)

type Objects []Object

// splitGrid split objects into integer grid.
func (os Objects) splitGrid() ([]int64, []int64, [][]bool) {
	var xSplit, ySplit []int64

	xUnique := make(map[int64]struct{})
	yUnique := make(map[int64]struct{})

	// remove duplicates
	for _, o := range os {
		if _, ok := xUnique[int64(o.X)]; !ok {
			xSplit = append(xSplit, int64(o.X))
			xUnique[int64(o.X)] = struct{}{}
		}

		if _, ok := xUnique[int64(o.X)+int64(o.Width)]; !ok {
			xSplit = append(xSplit, int64(o.X)+int64(o.Width))
			xUnique[int64(o.X)+int64(o.Width)] = struct{}{}
		}

		if _, ok := yUnique[int64(o.Y)]; !ok {
			ySplit = append(ySplit, int64(o.Y))
			yUnique[int64(o.Y)] = struct{}{}
		}

		if _, ok := yUnique[int64(o.Y)+int64(o.Height)]; !ok {
			ySplit = append(ySplit, int64(o.Y)+int64(o.Height))
			yUnique[int64(o.Y)+int64(o.Height)] = struct{}{}
		}
	}

	sort.Slice(xSplit, func(i, j int) bool { return xSplit[i] < xSplit[j] })
	sort.Slice(ySplit, func(i, j int) bool { return ySplit[i] < ySplit[j] })

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
			if x == int64(o.X) {
				xMin = int64(i)
			}

			if x == int64(o.X)+int64(o.Width) {
				xMax = int64(i)

				break
			}
		}

		// find y range iterations in grid
		var yMin, yMax int64

		for i, y := range ySplit {
			if y == int64(o.Y) {
				yMin = int64(i)
			}

			if y == int64(o.Y)+int64(o.Height) {
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
			X:      float64(xSplit[r.Origin.X]),
			Y:      float64(ySplit[r.Origin.Y]),
			Width:  float64(xSplit[r.Origin.X+int64(r.Width)] - xSplit[r.Origin.X]),
			Height: float64(ySplit[r.Origin.Y+int64(r.Height)] - ySplit[r.Origin.Y]),
		})
	}

	return result
}
