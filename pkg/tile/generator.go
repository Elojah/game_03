package tile

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/elojah/game_03/pkg/geometry"
)

// Params required for map generation.
type Params struct {
	Height          uint64
	Width           uint64
	CellHeight      uint64
	CellWidth       uint64
	WorldDensity    float64
	PlatformDensity float64
	PlatformPadding float64
	SizeMin         uint64
	SizeMax         uint64
	Distortion      uint64
	PathMin         uint64
	PathMax         uint64
	PathDistorsion  uint64

	Set Set
}

type Field int8

const (
	None Field = iota
	Void
	Ground
)

type GroundGenerator struct {
	Map [][]Field
}

// Display prints map in terminal.
// Debug only.
func (g GroundGenerator) Display() {
	var delimiter strings.Builder
	for i := 0; i < len(g.Map[0]); i++ {
		delimiter.WriteByte('_')
	}

	fmt.Println(delimiter.String())

	for i := 0; i < len(g.Map); i++ {
		fmt.Print("|")

		for j := 0; j < len(g.Map[i]); j++ {
			switch g.Map[i][j] {
			case None:
				fmt.Print(" ")
			case Void:
				fmt.Print(" ")
			case Ground:
				fmt.Print("x")
			default:
				fmt.Printf("%d", g.Map[i][j])
			}
		}
		fmt.Println("|")
	}
	fmt.Println(delimiter.String())
}

func (g *GroundGenerator) set(x, y int64, f Field) {
	if x >= 0 && x < int64(len(g.Map)) &&
		len(g.Map) > 0 &&
		y >= 0 && y < int64(len(g.Map[x])) {
		g.Map[x][y] = f
	}
}

func (g *GroundGenerator) get(x, y int64) (Field, bool) {
	if x >= 0 && x < int64(len(g.Map)) &&
		len(g.Map) > 0 &&
		y >= 0 && y < int64(len(g.Map[x])) {
		return g.Map[x][y], true
	}

	return None, false
}

func (g *GroundGenerator) Gen(params Params) {
	rand.Seed(time.Now().UnixNano())

	g.Map = make([][]Field, params.Height*params.CellHeight)

	for i := int64(0); i < int64(params.Height*params.CellHeight); i++ {
		g.Map[i] = make([]Field, params.Width*params.CellWidth)
	}

	nPlatforms := int64(float64(params.Height*params.CellHeight*params.Width*params.CellWidth) * params.WorldDensity)
	platforms := make([]geometry.Vec2, 0, nPlatforms)

	for i := 0; i < int(nPlatforms); i++ {
		platforms = append(platforms, geometry.Vec2{
			X: rand.Int63n(int64(params.Height * params.CellHeight)), //nolint: gosec
			Y: rand.Int63n(int64(params.Width * params.CellWidth)),   //nolint: gosec
		})
	}

	for _, p := range platforms {
		pl := g.generatePlatform(
			int64(params.SizeMin)+rand.Int63n(int64(params.SizeMax-params.SizeMin)), //nolint: gosec
			int64(params.SizeMin)+rand.Int63n(int64(params.SizeMax-params.SizeMin)), //nolint: gosec
			params.PlatformDensity,
			params.PlatformPadding,
		)
		for i := int64(0); i < int64(len(pl)); i++ {
			for j := int64(0); j < int64(len(pl[i])); j++ {
				// override only if current spot is empty
				if n, ok := g.get(i+p.X, j+p.Y); ok && n == None {
					g.set(i+p.X, j+p.Y, pl[i][j])
				}
			}
		}
	}
}

func (g GroundGenerator) generatePlatform(height int64, width int64, density float64, padding float64) [][]Field {
	p := make([][]Field, height)

	for i := int64(0); i < height; i++ {
		p[i] = make([]Field, width)

		// adjust width padding
		pad := (rand.Float64() * padding * float64(width)) - float64(width/2) //nolint: gosec, gomnd
		minW := math.Max(pad, 0)
		maxW := math.Min(float64(width), pad+float64(width))

		for j := int64(minW); j < int64(maxW); j++ {
			if n := rand.Float64(); n < density { //nolint: gosec
				p[i][j] = Ground
			} else {
				p[i][j] = Void
			}
		}
	}

	return p
}

func (g GroundGenerator) Tilemap(params Params) Map {
	m := NewMap()

	// Main layer
	layer := NewLayer()
	layer.ID = 1
	layer.Height = int(params.Height * params.CellHeight)
	layer.Width = int(params.Width * params.CellWidth)

	collideLayer := NewLayer()
	collideLayer.ID = 2
	collideLayer.Height = int(params.Height * params.CellHeight)
	collideLayer.Width = int(params.Width * params.CellWidth)
	collideLayer.Properties = append(collideLayer.Properties, Property{
		Name:  "collides",
		Value: true,
	})

	m.Height = int(params.Height * params.CellHeight)
	m.Width = int(params.Width * params.CellWidth)
	m.NextLayerID = 3
	m.NextObjectID = 1
	m.Tilesets = append(m.Tilesets, params.Set)

	if len(g.Map) == 0 {
		return m
	}

	size := len(g.Map) * len(g.Map[0])
	data := make([]byte, 0, 4*size)        //nolint: gomnd
	collideData := make([]byte, 0, 4*size) //nolint: gomnd

	for i := 0; i < len(g.Map); i++ {
		for j := 0; j < len(g.Map[i]); j++ {
			data = binary.LittleEndian.AppendUint32(data, uint32(g.Map[i][j]))
			collideData = binary.LittleEndian.AppendUint32(data, func(f Field) uint32 {
				if f == Void {
					return 1
				}

				return 0
			}(g.Map[i][j]))
		}
	}

	layer.Data = base64.StdEncoding.EncodeToString(data)
	collideLayer.Data = base64.StdEncoding.EncodeToString(collideData)

	m.Layers = append(m.Layers, layer, collideLayer)

	return m
}
