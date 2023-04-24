package mem

import (
	"context"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/rtc"
)

func (s *Store) Insert(ctx context.Context, pc rtc.PC) error {
	id := pc.ID.String()

	if _, ok := s.values[id]; ok {
		return errors.ErrConflict{Resource: "peer_connection", Index: id}
	}

	s.values[id] = pc

	return nil
}

func (s Store) Fetch(ctx context.Context, f rtc.Filter) (rtc.PC, error) {
	id := f.ID.String()

	result, ok := s.values[id]
	if !ok {
		return rtc.PC{}, errors.ErrNotFound{Resource: "peer_connection", Index: id}
	}

	return result, nil
}

func (s Store) Delete(ctx context.Context, f rtc.Filter) error {
	id := f.ID.String()

	if _, ok := s.values[id]; !ok {
		return errors.ErrNotFound{Resource: "peer_connection", Index: id}
	}

	delete(s.values, id)

	return nil
}
