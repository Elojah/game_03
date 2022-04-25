package room_test

import (
	"fmt"
	testing "testing"

	"github.com/elojah/game_03/pkg/room"
)

func TestNewCells(t *testing.T) {
	w := room.World{
		Height: 3,
		Width:  4,
	}

	cs := w.NewCells()
	for _, cl := range cs {
		for _, c := range cl {
			fmt.Printf("X:%d/Y:%d   -   ", c.X, c.Y)
		}

		fmt.Println("")
	}
}
