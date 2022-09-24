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
			CellHeight:      50,
			CellWidth:       120,
			WorldDensity:    0.02,
			PlatformDensity: 1,
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
