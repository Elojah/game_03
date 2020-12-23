package redis

import (
	"context"
	"fmt"

	"github.com/elojah/game_03/pkg/entity"
)

const entityKey = "entity:"

func (s *Store) Upsert(ctx context.Context, e entity.E) error {
	raw, err := e.Marshal()
	if err != nil {
		return err
	}

	if err := s.Set(ctx, entityKey+e.ID.String(), raw, 0).Err(); err != nil {
		return fmt.Errorf("upsert entity %s: %w", e.ID, err)
	}

	return nil
}

func (s *Store) FetchMany(ctx context.Context, f entity.Filter, callback func(entity.E) error) error {
	// val, err := s.Get(ctx, entityKey+id).Result()
	// if err != nil {
	// 	if !errors.Is(err, redis.Nil) {
	// 		return entity.E{}, fmt.Errorf("fetch entity %s: %w", id, err)
	// 	}

	// 	return entity.E{}, perrors.ErrNotFound{Resource: entityKey, Index: id}
	// }

	// var e entity.E
	// if err := e.Unmarshal([]byte(val)); err != nil {
	// 	return entity.E{}, fmt.Errorf("fetch entity %s: %w", id, err)
	// }

	// return e, nil
}

func (s *Store) Delete(ctx context.Context, f entity.Filter) error {

	// if err := s.Del(ctx, entityKey+id).Err(); err != nil {
	// 	return fmt.Errorf("remove entity %s: %w", id, err)
	// }

	return nil
}
