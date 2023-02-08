package tile_test

import (
	"testing"

	"github.com/elojah/game_03/pkg/tile"
)

func TestMinimumDissection(t *testing.T) {
	for _, d := range []struct {
		input  tile.Objects
		output tile.Objects
	}{
		{
			input: tile.Objects{
				{X: 1, Y: 1, Width: 5, Height: 3},
			},
			output: tile.Objects{
				{X: 1, Y: 1, Width: 5, Height: 3},
			},
		},
		{
			input: tile.Objects{
				{X: 1, Y: 1, Width: 5, Height: 3},
				{X: 2, Y: 2, Width: 1, Height: 1},
			},
			output: tile.Objects{
				{X: 1, Y: 1, Width: 5, Height: 3},
			},
		},
		{
			input: tile.Objects{
				{X: 1, Y: 1, Width: 5, Height: 3},
				{X: 6, Y: 1, Width: 3, Height: 3},
			},
			output: tile.Objects{
				{X: 1, Y: 1, Width: 8, Height: 3},
			},
		},
		{
			input: tile.Objects{
				{X: 1, Y: 1, Width: 5, Height: 3},
				{X: 6, Y: 1, Width: 3, Height: 3},
				{X: 1, Y: 4, Width: 8, Height: 4},
			},
			output: tile.Objects{
				{X: 1, Y: 1, Width: 8, Height: 7},
			},
		},
		{
			input: tile.Objects{
				{X: 1, Y: 1, Width: 5, Height: 3},
				{X: 6, Y: 1, Width: 3, Height: 3},
				{X: 1, Y: 4, Width: 10, Height: 4},
			},
			output: tile.Objects{
				{X: 1, Y: 1, Width: 8, Height: 7},
				{X: 9, Y: 4, Width: 2, Height: 4},
			},
		},
	} {
		result := d.input.MinimumDissection()

		if len(result) != len(d.output) {
			t.Errorf("result:%v\noutput:%v", result, d.output)

			continue
		}

		for i, o := range result {
			if o.X != d.output[i].X ||
				o.Y != d.output[i].Y ||
				o.Width != d.output[i].Width ||
				o.Height != d.output[i].Height {
				t.Errorf("\nexpected:%v\noutput:%v", o, d.output[i])

				continue
			}
		}
	}
}
