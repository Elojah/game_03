package wang

import (
	"encoding/base64"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/geometry"
	gtile "github.com/elojah/game_03/pkg/tile"
)

type Edge uint8

const (
	Top Edge = iota
	Right
	Down
	Left
)

type ID int

type Color byte

// string method for debug purpose.
func (cs Colors) string() string {
	b := strings.Builder{}
	b.WriteByte('_')

	for _, c := range cs {
		b.WriteString(strconv.Itoa(int(c)))
	}

	b.WriteByte('_')

	return b.String()
}

type Colors string

type Tile map[Edge]Colors

type Tiles map[ID]Tile

func (ts Tiles) rand() ID {
	for id := range ts {
		// return first iteration, using random map access iteration to fetch a random element
		return id
	}

	return 0
}

type orientedTile map[Colors]map[ID]struct{}

type orientedTiles map[Edge]orientedTile

type wangtiles []gtile.WangTile

func (ws wangtiles) oriented() (Tiles, orientedTiles) {
	ts := make(Tiles, len(ws))
	ot := orientedTiles{
		Top:   make(orientedTile),
		Right: make(orientedTile),
		Down:  make(orientedTile),
		Left:  make(orientedTile),
	}

	for _, wt := range ws {
		// discard wang tiles with missing information
		if len(wt.WangID) < 8 { //nolint: gomnd
			continue
		}

		t := Tile{
			Top:   Colors([]byte{wt.WangID[7], wt.WangID[0], wt.WangID[1]}), // Left to Right
			Right: Colors([]byte{wt.WangID[1], wt.WangID[2], wt.WangID[3]}), // Top to Down
			Down:  Colors([]byte{wt.WangID[5], wt.WangID[4], wt.WangID[3]}), // Left to Right
			Left:  Colors([]byte{wt.WangID[7], wt.WangID[6], wt.WangID[5]}), // Top to Down
		}

		// add tiles to general usage tiles struct
		ts[ID(wt.TileID)] = t

		// add tiles to oriented struct for fast access during populate
		if _, ok := ot[Top][t[Top]]; !ok {
			ot[Top][t[Top]] = map[ID]struct{}{ID(wt.TileID): {}}
		} else {
			ot[Top][t[Top]][ID(wt.TileID)] = struct{}{}
		}

		if _, ok := ot[Right][t[Right]]; !ok {
			ot[Right][t[Right]] = map[ID]struct{}{ID(wt.TileID): {}}
		} else {
			ot[Top][t[Top]][ID(wt.TileID)] = struct{}{}
		}

		if _, ok := ot[Down][t[Down]]; !ok {
			ot[Down][t[Down]] = map[ID]struct{}{ID(wt.TileID): {}}
		} else {
			ot[Down][t[Down]][ID(wt.TileID)] = struct{}{}
		}

		if _, ok := ot[Left][t[Left]]; !ok {
			ot[Left][t[Left]] = map[ID]struct{}{ID(wt.TileID): {}}
		} else {
			ot[Left][t[Left]][ID(wt.TileID)] = struct{}{}
		}
	}

	return ts, ot
}

type Grid [][]ID

func (g Grid) Tilemap(r geometry.Rect, ts gtile.Set, collisions map[int][]gtile.Object) (gtile.Map, error) {
	if r.Origin.Y < 0 || r.Origin.Y+int64(r.Height) > int64(len(g)) ||
		r.Origin.X < 0 || r.Origin.X+int64(r.Width) > int64(len(g[r.Origin.Y])) {
		return gtile.Map{}, errors.ErrInvalidDimension{}
	}

	m := gtile.NewMap()

	m.Height = int(r.Height)
	m.Width = int(r.Width)
	m.Tilesets = append(m.Tilesets, ts.CleanCopy())

	// main layer
	layer := gtile.NewLayer()
	layer.ID = 1
	layer.Height = m.Height
	layer.Width = m.Width

	// collision layer
	clayer := gtile.NewLayer()
	clayer.ID = 2
	clayer.Type = "objectgroup"
	clayer.Properties = append(clayer.Properties, gtile.Property{Name: "collision", Type: "bool", Value: true})
	clayer.Visible = false
	clayer.Height = m.Height
	clayer.Width = m.Width

	size := layer.Height * layer.Width
	data := make([]byte, 0, 4*size) //nolint: gomnd

	var ii, jj int

	for i := r.Origin.Y; i < r.Origin.Y+int64(r.Height); i++ {
		jj = 0

		for j := r.Origin.X; j < r.Origin.X+int64(r.Width); j++ {
			data = binary.LittleEndian.AppendUint32(data, uint32(g[i][j]))

			// add collide object to clayer
			if objects, ok := collisions[int(g[i][j])]; ok {
				for _, o := range objects {
					o := o
					o.X += float64(jj * ts.TileWidth)
					o.Y += float64(ii * ts.TileHeight)
					clayer.Objects = append(clayer.Objects, o)
				}
			}
			jj++
		}
		ii++
	}

	layer.Data = base64.StdEncoding.EncodeToString(data)

	m.Layers = append(m.Layers, layer)
	m.NextLayerID = 2

	if len(clayer.Objects) > 0 {
		clayer.Objects = gtile.Objects(clayer.Objects).MinimumDissection()

		m.Layers = append(m.Layers, clayer)
		m.NextLayerID = 3
	}

	return m, nil
}

type Heuristic func(candidates map[ID]struct{}, x int64, y int64, ts Tiles) map[ID]struct{}

// GenerateFlat is a debug function to display flat tileset.
func (g *Grid) GenerateFlat(w gtile.WangSet, height int64, width int64) {
	result := make(Grid, height)
	for i := range result {
		result[i] = make([]ID, width)
	}

	for i := int64(0); i < height; i++ {
		for j := int64(0); j < width; j++ {
			result[i][j] = ID(((i * height) + width + 1) % 504)
		}
	}

	*g = result
}

func (g *Grid) Generate(w gtile.WangSet, height int64, width int64, hs ...Heuristic) { //nolint: gocognit
	ts, ot := wangtiles(w.WangTiles).oriented()

	result := make(Grid, height)
	for i := range result {
		result[i] = make([]ID, width)
	}

	for i := int64(0); i < height; i++ {
		for j := int64(0); j < width; j++ {
			constraints := make(map[Edge]Colors)

			// set only Top & Left constraints for usual Top Left run
			if i > 0 {
				t, ok := ts[result[i-1][j]]

				if !ok {
					// should never happen, it means we set an invalid id tile in a previous loop
					continue
				}

				constraints[Top] = t[Down]
			}

			if j > 0 {
				t, ok := ts[result[i][j-1]]

				if !ok {
					// should never happen, it means we set an invalid id tile in a previous loop
					continue
				}

				constraints[Left] = t[Right]
			}

			if len(constraints) == 0 {
				// no constraints (first tile to put), put random one
				result[i][j] = ts.rand()

				continue
			}

			var candidates map[ID]struct{}

			filter := false

			for o, constraint := range constraints {
				c, ok := ot[o][constraint]
				if !ok {
					// constraint color not found, return invalid index
					// TODO: error case ? ignore and set 0 ? warning ?
					// fmt.Println("single constraint not found")

					result[i][j] = 0

					break
				}

				if !filter {
					// first pass, add all ids candidates
					candidates = c
					filter = true
				} else {
					// second pass, if id also exists in candidates, add as solution
					tmp := make(map[ID]struct{})

					for k := range c {
						if _, ok := candidates[k]; ok {
							tmp[k] = struct{}{}
						}
					}

					candidates = tmp
				}
			}

			if len(candidates) == 0 {
				// constraint color not found, return invalid index
				// TODO: error case ? ignore and set 0 ? warning ?
				// fmt.Println("single constraint not found")

				result[i][j] = 0

				break
			}

			for _, h := range hs {
				candidates = h(candidates, i, j, ts)
			}

			for c := range candidates {
				// pick first result (assign and break loop)
				result[i][j] = c
				break
			}
		}
	}

	*g = result
}
