package mem

import "github.com/elojah/game_03/pkg/rtc"

type Store struct {
	values map[string]rtc.PC
}

func NewStore() Store {
	return Store{
		values: make(map[string]rtc.PC),
	}
}
