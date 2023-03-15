package template

import "github.com/elojah/game_03/pkg/tile/wang"

func AttunedHeuristic(candidates map[wang.ID]struct{}, y int64, x int64, ts wang.Tiles) map[wang.ID]struct{} {
	result := map[wang.ID]struct{}{}
	maxAttuned := 0

	for c := range candidates {
		occurences := make(map[byte]int)
		for _, w := range []byte{
			ts[c][wang.Top][0],   // top left
			ts[c][wang.Top][1],   // top
			ts[c][wang.Top][2],   // top right
			ts[c][wang.Left][1],  // left
			ts[c][wang.Right][1], // right
			ts[c][wang.Left][2],  // down left
			ts[c][wang.Down][1],  // down
			ts[c][wang.Down][2],  // down right
		} {
			if n, ok := occurences[w]; ok {
				occurences[w] = n + 1
			} else {
				occurences[w] = 1
			}
		}

		for _, n := range occurences {
			if n > maxAttuned {
				maxAttuned = n
				result = map[wang.ID]struct{}{c: {}}
			} else if n == maxAttuned {
				result[c] = struct{}{}
			}
		}

	}

	return result
}
