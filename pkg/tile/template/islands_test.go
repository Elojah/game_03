package template_test

import (
	"testing"

	"github.com/elojah/game_03/pkg/tile/template"
)

func TestNewIslands(t *testing.T) {
	for _, d := range []struct {
		name string
		opts template.IslandsOptions
	}{
		{
			name: "basic",
			opts: template.IslandsOptions{
				Height:        80,
				Width:         150,
				IslandHeight:  8,
				IslandWidth:   7,
				PaddingHeight: 4,
				PaddingWidth:  4,
				PathWidth:     4,
			},
		},
	} {
		t := template.NewIslands(d.opts)

		_ = t
	}
}
