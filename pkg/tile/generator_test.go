package tile_test

import (
	"testing"

	"github.com/elojah/game_03/pkg/tile"
)

func TestGenerator(t *testing.T) {
	for _, p := range []tile.Params{
		{
			Height:          1,
			Width:           1,
			CellHeight:      20,
			CellWidth:       15,
			WorldDensity:    0.01,
			PlatformDensity: 1,
			PlatformPadding: 0.8,
			SizeMin:         5,
			SizeMax:         10,
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
