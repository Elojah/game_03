package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type FilterWorld struct {
	ID  ulid.ID
	IDs []ulid.ID

	All bool

	State []byte
	Size  int
}

type StoreWorld interface {
	InsertWorld(context.Context, World) error
	FetchWorld(context.Context, FilterWorld) (World, error)
	DeleteWorld(context.Context, FilterWorld) error
	FetchManyWorld(context.Context, FilterWorld) ([]World, []byte, error)
}

func (w World) NewCells() [][]Cell { //nolint: gocognit
	// Create cells with ID
	cells := make([][]Cell, w.Height)
	for i := range cells {
		cells[i] = make([]Cell, w.Width)
		for j := range cells[i] {
			cells[i][j] = Cell{
				ID:         ulid.NewID(),
				Contiguous: make(map[int32]ulid.ID),
			}
		}
	}

	// Set contiguous IDs
	for i := range cells {
		for j := range cells[i] {
			// up
			if i != 0 {
				cells[i][j].Contiguous[int32(Up)] = cells[i-1][j].ID
			}

			// up right
			if i != 0 && j != len(cells[i])-1 {
				cells[i][j].Contiguous[int32(UpRight)] = cells[i-1][j+1].ID
			}

			// right
			if j != len(cells[i])-1 {
				cells[i][j].Contiguous[int32(Right)] = cells[i][j+1].ID
			}

			// down right
			if i != len(cells)-1 && j != len(cells[i])-1 {
				cells[i][j].Contiguous[int32(DownRight)] = cells[i+1][j+1].ID
			}

			// down
			if i != len(cells)-1 {
				cells[i][j].Contiguous[int32(Down)] = cells[i+1][j].ID
			}

			// down left
			if i != len(cells)-1 && j != 0 {
				cells[i][j].Contiguous[int32(DownLeft)] = cells[i+1][j-1].ID
			}

			// left
			if j != 0 {
				cells[i][j].Contiguous[int32(Left)] = cells[i][j-1].ID
			}

			// up left
			if i != 0 && j != 0 {
				cells[i][j].Contiguous[int32(UpLeft)] = cells[i-1][j-1].ID
			}
		}
	}

	return cells
}
