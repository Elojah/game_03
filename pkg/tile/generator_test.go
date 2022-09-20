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
			CellHeight:      200,
			CellWidth:       200,
			WorldDensity:    0.05,
			PlatformDensity: 0.1,
			PlatformPadding: 0.9,
			SizeMin:         5,
			SizeMax:         20,
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
