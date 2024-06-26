package redis

import (
	"context"
	"fmt"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/redis/rueidis"
)

const (
	key = "entity:"
)

// InsertCache implemented with redis.
func (c *Cache) InsertCache(ctx context.Context, e entity.E) error {
	raw, err := e.Marshal()
	if err != nil {
		return fmt.Errorf("insert entity %s at %d: %w", e.ID.String(), e.At, err)
	}

	if err := c.Do(
		ctx,
		c.B().Zadd().Key(key+e.ID.String()).Nx().ScoreMember().ScoreMember(
			float64(e.At),
			rueidis.BinaryString(raw),
		).Build(),
	).Error(); err != nil {
		return fmt.Errorf("insert entity %s at %d: %w", e.ID.String(), e.At, err)
	}

	return nil
}

// FetchManyCache implemented with redis.
func (c *Cache) FetchManyCache(ctx context.Context, filter entity.FilterCache) ([]entity.E, error) {
	var cmd rueidis.Completed
	if filter.Reverse {
		cmd = c.B().Zrevrangebyscore().
			Key(key+filter.ID.String()).
			Max(filter.Max).
			Min(filter.Min).
			Limit(0, filter.Size).
			Build()
	} else {
		cmd = c.B().Zrangebyscore().
			Key(key+filter.ID.String()).
			Min(filter.Min).
			Max(filter.Max).
			Limit(0, filter.Size).
			Build()
	}

	vals, err := c.Do(ctx, cmd).AsStrSlice()
	if err != nil {
		return nil, fmt.Errorf("fetch entity %s:%w", filter.ID.String(), err)
	}

	es := make([]entity.E, len(vals))
	for i := range vals {
		if err := es[i].Unmarshal([]byte(vals[i])); err != nil {
			return nil, fmt.Errorf("fetch entity for entity %s:%w", filter.ID.String(), err)
		}
	}

	return es, nil
}

// DeleteCache implemented with redis.
func (c *Cache) DeleteCache(ctx context.Context, filter entity.FilterCache) error {
	if err := c.Do(
		ctx,
		c.B().Zremrangebyscore().Key(key+filter.ID.String()).Min(filter.Min).Max(filter.Max).Build(),
	).Error(); err != nil {
		return fmt.Errorf("remove entity %s:%w", filter.ID.String(), err)
	}

	return nil
}
