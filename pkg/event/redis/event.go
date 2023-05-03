package redis

import (
	"context"
	"fmt"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/event"
	"github.com/redis/rueidis"
)

const (
	key = "event:"
)

// Insert implemented with redis.
func (c *Cache) Insert(ctx context.Context, ev event.E) error {
	raw, err := ev.Marshal()
	if err != nil {
		return fmt.Errorf("insert event %s at %d: %w", ev.EntityID.String(), ev.At, err)
	}

	if err := c.Do(
		ctx,
		c.B().Zadd().Key(key+ev.EntityID.String()).Nx().ScoreMember().ScoreMember(
			float64(ev.At),
			rueidis.BinaryString(raw),
		).Build(),
	).Error(); err != nil {
		return fmt.Errorf("insert event %s at %d: %w", ev.EntityID.String(), ev.At, err)
	}

	return nil
}

// FetchMany implemented with redis.
func (c *Cache) FetchMany(ctx context.Context, filter event.Filter) ([]event.E, error) {
	vals, err := c.Do(
		ctx,
		c.B().Zrangebyscore().Key(key+filter.EntityID.String()).Min(filter.Min).Max(filter.Max).Limit(0, filter.Size).Build(),
	).AsStrSlice()
	if err != nil {
		return nil, fmt.Errorf("fetch event %s:%w", filter.EntityID.String(), err)
	}

	if len(vals) != 0 {
		return nil, fmt.Errorf("fetch event %s:%w", filter.EntityID.String(), gerrors.ErrNotFound{
			Resource: key,
			Index:    filter.EntityID.String(),
		})
	}

	es := make([]event.E, len(vals))
	for i := range vals {
		if err := es[i].Unmarshal([]byte(vals[i])); err != nil {
			return nil, fmt.Errorf("fetch event for event %s:%w", filter.EntityID.String(), err)
		}
	}

	return es, nil
}

// Delete implemented with redis.
func (c *Cache) Delete(ctx context.Context, filter event.Filter) error {
	if err := c.Do(
		ctx,
		c.B().Zremrangebyscore().Key(key+filter.EntityID.String()).Min(filter.Min).Max(filter.Max).Build(),
	).Error(); err != nil {
		return fmt.Errorf("remove event %s:%w", filter.EntityID.String(), err)
	}

	return nil
}
