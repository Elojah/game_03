package tile

import (
	"fmt"
	"strings"
	"testing"
)

func (wg wangtiles) print() {
	b := strings.Builder{}

	for _, line := range wg {
		for _, id := range line {
			// top
			if id&(1<<7) > 0 {
				b.WriteRune('#')
			} else {
				b.WriteRune('_')
			}

			// top-right
			if id&(1<<6) > 0 {
				b.WriteRune('#')
			} else {
				b.WriteRune('_')
			}

			// right
			if id&(1<<5) > 0 {
				b.WriteRune('#')
			} else {
				b.WriteRune('_')
			}

			// bottom-right
			if id&(1<<4) > 0 {
				b.WriteRune('#')
			} else {
				b.WriteRune('_')
			}

			// bottom
			if id&(1<<3) > 0 {
				b.WriteRune('#')
			} else {
				b.WriteRune('_')
			}

			// bottom-left
			if id&(1<<2) > 0 {
				b.WriteRune('#')
			} else {
				b.WriteRune('_')
			}

			// left
			if id&(1<<1) > 0 {
				b.WriteRune('#')
			} else {
				b.WriteRune('_')
			}

			// top-left
			if id&(1<<0) > 0 {
				b.WriteRune('#')
			} else {
				b.WriteRune('_')
			}

			b.WriteRune(' ')
		}

		b.WriteRune('\n')
	}

	fmt.Print(b.String())
}

func TestNewWangtiles(t *testing.T) {
	t.Run("regular", func(t *testing.T) {
		wg, err := NewWangtiles(10, 10, Corner)
		if err != nil {
			t.Error(err)

			return
		}

		_ = wg
		// wg.print()
	})
}
