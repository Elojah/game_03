package tile

type Orientation uint8

const (
	Up = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

type T struct {
	Tile

	mask byte
}

func (t *T) SetContiguous(os ...Orientation) {
	for _, o := range os {
		t.mask |= (1 << o)
	}
}
