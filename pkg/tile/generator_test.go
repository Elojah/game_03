package tile_test

import (
	"testing"

	"github.com/elojah/game_03/pkg/tile"
)

func TestGenerator(t *testing.T) {
	for _, p := range []tile.Params{
		{
			Height:          10,
			Width:           50,
			CellHeight:      10,
			CellWidth:       10,
			WorldDensity:    0.9,
			PlatformDensity: 0.7,
			PlatformPadding: 0.9,
			SizeMin:         5,
			SizeMax:         15,
		},
	} {
		t.Run(
			"display",
			func(t *testing.T) {
				g := tile.GroundGenerator{}
				g.Gen(p)
				g.Display()
			},
		)
	}
}
