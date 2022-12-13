package tile

import (
	"encoding/base64"
	"encoding/binary"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/geometry"
)

type WangType string

const (
	Corner = "corner"
	Edge   = "edge"
	All    = "all"
)

// wangid represents a wang tile with bytes set in order:
// top, top-right, right, bottom-right, bottom, bottom-left, left, top-left.
// e.g: 01001010.
type wangid uint8

// wangtiles represents wangtiles id for tilemap data generation.
type wangtiles [][]wangid

func newWangID(wang string) wangid {
	result := wangid(0)

	// don't support wang tiles with edge/corner indications != 8
	if len(wang) != 8 { //nolint: gomnd
		return result
	}

	for i, b := range []byte(wang) {
		if b == 1 {
			result |= (1 << (7 - i)) //nolint: gomnd
		}
	}

	return result
}

func newWangIDsMap(wts []WangTile) map[wangid]int {
	result := make(map[wangid]int)

	for _, wt := range wts {
		result[newWangID(string(wt.WangID))] = wt.TileID
	}

	return result
}

// wangid returns wangid with set byte.
// order: top, top-right, right, bottom-right, bottom, bottom-left, left, top-left.
func (g grid) wangid(i int, j int) wangid {
	result := wangid(0)
	i = (2 * i) + 1 //nolint: gomnd
	j = (2 * j) + 1 //nolint: gomnd

	// overflow check
	if i < 0 || i > len(g) || j < 0 || j > len(g[i]) {
		return 0
	}

	// top
	if i > 0 && g[i-1][j] {
		result |= (1 << 7) //nolint: gomnd
	}

	// top-right
	if i > 0 && j < len(g[i]) && g[i-1][j+1] {
		result |= (1 << 6) //nolint: gomnd
	}

	// right
	if j < len(g[i]) && g[i][j+1] {
		result |= (1 << 5) //nolint: gomnd
	}

	// bottom-right
	if i < len(g) && j < len(g[i]) && g[i+1][j+1] {
		result |= (1 << 4) //nolint: gomnd
	}

	// bottom
	if i < len(g) && g[i+1][j] {
		result |= (1 << 3) //nolint: gomnd
	}

	// bottom-left
	if i < len(g) && j > 0 && g[i+1][j-1] {
		result |= (1 << 2) //nolint: gomnd
	}

	// left
	if j > 0 && g[i][j-1] {
		result |= (1 << 1)
	}

	// top-left
	if i > 0 && j > 0 && g[i-1][j-1] {
		result |= (1 << 0)
	}

	return result
}

// NewWangtiles generates a wangid wangtiles.
func NewWangtiles(height int64, width int64, t WangType) (wangtiles, error) {
	g := newRandGrid(height, width, t)

	wg := make(wangtiles, height)
	for i := range wg {
		wg[i] = make([]wangid, width)
	}

	for i, line := range wg {
		for j := range line {
			line[j] = g.wangid(i, j)
		}
	}

	return wg, nil
}

func (wg wangtiles) Tilemap(r geometry.Rect, ts Set) (Map, error) {
	if len(wg) == 0 {
		return Map{}, nil
	}

	if len(ts.WangSets) == 0 {
		return Map{}, errors.ErrMissingWangSet{ID: ts.ID.String()}
	}

	if r.Origin.Y < 0 || r.Origin.Y+int64(r.Height) > int64(len(wg)) ||
		r.Origin.X < 0 || r.Origin.X+int64(r.Width) > int64(len(wg[r.Origin.Y])) {
		return Map{}, errors.ErrInvalidDimension{}
	}

	m := NewMap()

	m.Height = int(r.Height)
	m.Width = int(r.Width)
	m.NextLayerID = 2
	m.NextObjectID = 1
	m.Tilesets = append(m.Tilesets, ts.CleanCopy())

	// Main layer
	layer := NewLayer()
	layer.ID = 1
	layer.Height = m.Height
	layer.Width = m.Width

	size := layer.Height * layer.Width
	data := make([]byte, 0, 4*size) //nolint: gomnd

	// TODO: Always fetch first ? Use property instead ?
	wm := newWangIDsMap(ts.WangSets[0].WangTiles)

	for i := r.Origin.Y; i < r.Origin.Y+int64(r.Height); i++ {
		for j := r.Origin.X; j < r.Origin.X+int64(r.Width); j++ {
			id, ok := wm[wg[i][j]]

			if !ok {
				// return Map{}, errors.ErrMissingWangSet{ID: ts.ID.String()}
				id = 0
			}

			// collideData = binary.LittleEndian.AppendUint32(collideData, func(f Field) wangid {
			// 	if f.Collides() {
			// 		return 1
			// 	}

			// 	return 0
			// }(f))

			data = binary.LittleEndian.AppendUint32(data, uint32(id))
		}
	}

	layer.Data = base64.StdEncoding.EncodeToString(data)

	m.Layers = append(m.Layers, layer)

	return m, nil
}
