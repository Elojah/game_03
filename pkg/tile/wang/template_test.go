package wang_test

import (
	"fmt"
	"testing"

	gtile "github.com/elojah/game_03/pkg/tile"
	"github.com/elojah/game_03/pkg/tile/wang"
)

func TestNewTemplate(t *testing.T) {
	for _, d := range []struct {
		name   string
		ws     gtile.WangSet
		height int64
		width  int64
	}{
		{
			name:   "basic",
			height: 100,
			width:  70,
			ws: gtile.WangSet{
				WangTiles: []gtile.WangTile{
					{
						TileID: 101,
						WangID: []byte{11, 14, 13, 0, 12, 0, 0, 0},
					},
					{
						TileID: 102,
						WangID: []byte{12, 0, 13, 0, 11, 14, 13, 0},
					},
					{
						TileID: 103,
						WangID: []byte{12, 0, 13, 0, 12, 14, 0, 14},
					},
					{
						TileID: 104,
						WangID: []byte{11, 0, 0, 0, 11, 14, 13, 0},
					},
				},
			},
		},
		{
			name:   "complex",
			height: 100,
			width:  70,
			ws: gtile.WangSet{
				WangTiles: []gtile.WangTile{
					{
						TileID: 38,
						WangID: []byte{0, 0, 3, 1, 3, 0, 0, 0},
					},
					{
						TileID: 39,
						WangID: []byte{0, 0, 3, 1, 1, 1, 3, 0},
					},
					{
						TileID: 40,
						WangID: []byte{0, 0, 0, 0, 3, 1, 3, 0},
					},
					{
						TileID: 42,
						WangID: []byte{2, 2, 2, 0, 3, 1, 3, 2},
					},
					{
						TileID: 43,
						WangID: []byte{2, 2, 3, 1, 3, 0, 2, 2},
					},
					{
						TileID: 44,
						WangID: []byte{2, 0, 0, 0, 3, 1, 3, 2},
					},
					{
						TileID: 45,
						WangID: []byte{2, 0, 0, 0, 3, 1, 3, 2},
					},
					{
						TileID: 73,
						WangID: []byte{0, 0, 3, 1, 3, 0, 0, 0},
					},
					{
						TileID: 74,
						WangID: []byte{3, 1, 1, 1, 1, 1, 3, 0},
					},
					{
						TileID: 75,
						WangID: []byte{1, 1, 1, 1, 1, 1, 1, 1},
					},
					{
						TileID: 76,
						WangID: []byte{3, 0, 3, 1, 1, 1, 1, 1},
					},
					{
						TileID: 77,
						WangID: []byte{0, 0, 0, 0, 3, 1, 3, 0},
					},
					{
						TileID: 78,
						WangID: []byte{2, 2, 2, 0, 3, 1, 3, 2},
					},
					{
						TileID: 79,
						WangID: []byte{2, 2, 3, 1, 3, 0, 2, 2},
					},
					{
						TileID: 80,
						WangID: []byte{2, 2, 3, 1, 3, 0, 0, 0},
					},
					{
						TileID: 81,
						WangID: []byte{2, 2, 3, 1, 3, 0, 0, 0},
					},
					{
						TileID: 109,
						WangID: []byte{3, 1, 1, 1, 3, 0, 0, 0},
					},
					{
						TileID: 110,
						WangID: []byte{1, 1, 1, 1, 1, 1, 1, 1},
					},
					{
						TileID: 111,
						WangID: []byte{1, 1, 1, 1, 1, 1, 1, 1},
					},
					{
						TileID: 112,
						WangID: []byte{1, 1, 1, 1, 1, 1, 1, 1},
					},
					{
						TileID: 113,
						WangID: []byte{3, 0, 0, 0, 3, 1, 1, 1},
					},
					{
						TileID: 114,
						WangID: []byte{3, 1, 1, 1, 1, 1, 3, 2},
					},
					{
						TileID: 115,
						WangID: []byte{3, 2, 3, 1, 1, 1, 1, 1},
					},
					{
						TileID: 116,
						WangID: []byte{2, 2, 3, 1, 1, 1, 3, 2},
					},
					{
						TileID: 145,
						WangID: []byte{3, 1, 3, 2, 2, 0, 0, 0},
					},
					{
						TileID: 146,
						WangID: []byte{1, 1, 1, 1, 3, 2, 3, 1},
					},
					{
						TileID: 147,
						WangID: []byte{1, 1, 1, 1, 1, 1, 1, 1},
					},
					{
						TileID: 148,
						WangID: []byte{1, 1, 3, 2, 3, 1, 1, 1},
					},
					{
						TileID: 149,
						WangID: []byte{3, 0, 0, 0, 2, 2, 3, 1},
					},
					{
						TileID: 150,
						WangID: []byte{2, 0, 3, 1, 1, 1, 3, 2},
					},
					{
						TileID: 151,
						WangID: []byte{2, 2, 3, 1, 1, 1, 3, 0},
					},
					{
						TileID: 152,
						WangID: []byte{2, 2, 2, 0, 3, 1, 3, 0},
					},
					{
						TileID: 153,
						WangID: []byte{2, 2, 2, 0, 3, 1, 3, 0},
					},
					{
						TileID: 181,
						WangID: []byte{2, 2, 2, 0, 0, 0, 0, 0},
					},
					{
						TileID: 182,
						WangID: []byte{3, 1, 3, 2, 2, 0, 2, 2},
					},
					{
						TileID: 183,
						WangID: []byte{1, 1, 3, 2, 2, 2, 3, 1},
					},
					{
						TileID: 184,
						WangID: []byte{3, 2, 2, 0, 2, 2, 3, 1},
					},
					{
						TileID: 185,
						WangID: []byte{2, 0, 0, 0, 0, 0, 2, 2},
					},
					{
						TileID: 186,
						WangID: []byte{1, 1, 1, 1, 3, 2, 3, 1},
					},
					{
						TileID: 187,
						WangID: []byte{1, 1, 3, 2, 3, 1, 1, 1},
					},
					{
						TileID: 188,
						WangID: []byte{2, 0, 3, 1, 3, 0, 2, 2},
					},
					{
						TileID: 189,
						WangID: []byte{2, 0, 3, 1, 3, 0, 2, 2},
					},
					{
						TileID: 218,
						WangID: []byte{2, 2, 2, 0, 0, 0, 0, 0},
					},
					{
						TileID: 219,
						WangID: []byte{2, 2, 2, 0, 0, 0, 2, 2},
					},
					{
						TileID: 220,
						WangID: []byte{2, 0, 0, 0, 0, 0, 2, 2},
					},
					{
						TileID: 222,
						WangID: []byte{3, 1, 1, 1, 3, 0, 2, 2},
					},
					{
						TileID: 223,
						WangID: []byte{3, 2, 2, 0, 3, 1, 1, 1},
					},
					{
						TileID: 224,
						WangID: []byte{2, 2, 2, 0, 3, 1, 3, 0},
					},
					{
						TileID: 225,
						WangID: []byte{2, 2, 2, 0, 3, 1, 3, 0},
					},
					{
						TileID: 258,
						WangID: []byte{2, 0, 3, 1, 1, 1, 3, 2},
					},
					{
						TileID: 259,
						WangID: []byte{2, 2, 3, 1, 1, 1, 3, 0},
					},
					{
						TileID: 260,
						WangID: []byte{2, 0, 3, 1, 3, 0, 2, 2},
					},
					{
						TileID: 261,
						WangID: []byte{2, 0, 3, 1, 3, 0, 2, 2},
					},
				},
			},
		},
	} {
		t := wang.NewTemplate(d.ws, d.height, d.width)

		fmt.Println(t.String())
	}
}
