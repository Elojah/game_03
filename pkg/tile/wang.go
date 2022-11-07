package tile

import (
	"encoding/binary"
	"math/rand"
)

// wangtiles represents wangtiles id for tilemap data generation.
type wangtiles [][]uint8

// grid represents underlying generation to create correct wangtiles map.
type grid [][]bool

// bit structure for fast random node generation (1 or 0).
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

func newGrid(height int, width int) grid {
	b := bit{}

	g := make(grid, height) //nolint: gomnd
	for i := range g {
		g[i] = make([]bool, width) //nolint: gomnd

		for j := range g[i] {
			g[i][j] = b.rand()
		}
	}

	return g
}

// wangid returns uint8 with set byte.
// order: top, top-right, right, bottom-right, bottom, bottom-left, left, top-left.
func (g grid) wangid(i int, j int) uint8 {
	result := uint8(0)

	// overflow check
	if i < 0 || i > len(g) || j < 0 || j > len(g[i]) {
		return 0
	}

	// top
	if i > 0 && g[i-1][j] {
		result &= (1 << 7)
	}

	// top-right
	if i > 0 && j < len(g[i]) && g[i-1][j+1] {
		result &= (1 << 6)
	}

	// right
	if j < len(g[i]) && g[i][j+1] {
		result &= (1 << 5)
	}

	// bottom-right
	if i < len(g) && j < len(g[i]) && g[i+1][j+1] {
		result &= (1 << 4)
	}

	// bottom
	if i < len(g) && g[i+1][j] {
		result &= (1 << 3)
	}

	// bottom-left
	if i < len(g) && j > 0 && g[i+1][j-1] {
		result &= (1 << 2)
	}

	// left
	if j > 0 && g[i][j-1] {
		result &= (1 << 1)
	}

	// top-left
	if i > 0 && j > 0 && g[i-1][j-1] {
		result &= (1 << 0)
	}

	return result
}

// newWangtiles generates a wangid wangtiles.
func newWangtiles(height int, width int) (wangtiles, error) {
	g := newGrid((2*height)+1, (2*width)+1)

	wg := make(wangtiles, height)
	for i := range wg {
		wg[i] = make([]uint8, width)
	}

	for _, line := range wg {
		for j := range line {
			// line[j] =
		}
	}

	return wg, nil
}

func (wg wangtiles) layers() []Layer {
	if len(wg) == 0 {
		return nil
	}

	// Main layer
	layer := NewLayer()
	layer.ID = 1
	layer.Height = len(wg)
	layer.Width = len(wg[0])

	size := len(wg) * len(wg[0])
	data := make([]byte, 0, 4*size) //nolint: gomnd

	for _, line := range wg {
		for _, id := range line {
			data = binary.LittleEndian.AppendUint32(data, id)
			// collideData = binary.LittleEndian.AppendUint32(collideData, func(f Field) uint8 {
			// 	if f.Collides() {
			// 		return 1
			// 	}

			// 	return 0
			// }(f))
		}
	}

	return []Layer{layer}
}
