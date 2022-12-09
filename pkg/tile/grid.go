package tile

import (
	"math/rand"
	"time"

	"github.com/elojah/game_03/pkg/geometry"
)

// grid represents underlying generation structure to create correct wangtiles map.
type grid [][]bool

// bit structure for fast random bool generation (1 or 0).
type bit struct {
	n         int64
	remaining int64
}

// rand return 1 or 0.
func (b *bit) rand() bool {
	if b.remaining <= 0 {
		b.n = rand.Int63() //nolint: gosec
		b.remaining = 63
	}

	b.remaining--

	if b.n&(1<<(b.remaining+1)) > 0 {
		return true
	}

	return false
}

func newRandGrid(height int64, width int64, t WangType) grid {
	rand.Seed(time.Now().UnixNano())

	b := bit{}

	g := make(grid, (2*height)+1) //nolint: gomnd
	for i := range g {
		g[i] = make([]bool, (2*width)+1) //nolint: gomnd

		for j := range g[i] {
			if (t == Corner && (i%2 == 1 || j%2 == 1)) ||
				(t == Edge && (i%2 == 0 || j%2 == 0)) {
				// force unwanted nodes to be set false
				g[i][j] = false
			} else if i == 0 || j == 0 ||
				i == len(g)-1 || j == len(g[i])-1 {
				// force map borders to be set true
				g[i][j] = true
			} else {
				// set random node
				g[i][j] = b.rand()
			}
		}
	}

	return g
}

type generator struct {
	Grid grid

	PlatformDensity float64

	PlatformHeightMin      int
	PlatformHeightVariance int

	PlatformWidthMin      int
	PlatformWidthVariance int
	PlatformWidthPadding  int
}

func (g *generator) set(i int, j int) {
	if (i >= 0 && i < len(g.Grid)) &&
		(j >= 0 && j < len(g.Grid[i])) {
		g.Grid[i][j] = true
	}
}

func (g *generator) Generate(height int64, width int64, t WangType) error {
	gr := make(grid, (2*height)+1) //nolint: gomnd
	for i := range gr {
		gr[i] = make([]bool, (2*width)+1) //nolint: gomnd
	}

	rand.Seed(time.Now().UnixNano())

	nPlatforms := int64(float64(height*width) * g.PlatformDensity)

	platforms := make([]geometry.Vec2, 0, (nPlatforms))

	for i := 0; i < int(nPlatforms); i++ {
		platforms = append(platforms, geometry.Vec2{
			X: rand.Int63n(width),  //nolint: gosec
			Y: rand.Int63n(height), //nolint: gosec
		})
	}

	for _, p := range platforms {
		h := g.PlatformHeightMin * rand.Intn(g.PlatformHeightVariance) //nolint: gosec
		w := g.PlatformWidthMin * rand.Intn(g.PlatformWidthVariance)   //nolint: gosec

		for i := 0; i < h; i++ {
			pad := rand.Intn(g.PlatformWidthPadding) //nolint: gosec

			for j := 0; j < w; j++ {
				if j < pad || j > w-pad {
					continue
				}

				g.set(int(p.X)+i, int(p.Y)+j)
			}
		}
	}

	return nil
}
