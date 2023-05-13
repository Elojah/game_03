package redis

import (
	"context"
	"fmt"

	"github.com/elojah/game_03/pkg/ability"
	"github.com/redis/rueidis"
)

const (
	key = "ability:"
)

// Insert implementation for ability in redis.
func (c *Cache) Insert(ctx context.Context, ab ability.A) error {
	raw, err := ab.Marshal()
	if err != nil {
		return err
	}

	if err := c.Do(
		ctx,
		c.B().Set().Key(key+ab.ID.String()).Value(rueidis.BinaryString(raw)).Build(),
	).Error(); err != nil {
		return fmt.Errorf("upsert ability %s: %w", ab.ID.String(), err)
	}

	return nil
}

// Fetch implementation for ability in redis.
func (c *Cache) Fetch(ctx context.Context, filter ability.Filter) (ability.A, error) {
	val, err := c.DoCache(ctx, c.B().Get().Key(key+filter.ID.String()).Cache(), 0).AsBytes()
	if err != nil {
		return ability.A{}, fmt.Errorf("fetch ability %s: %w", filter.ID.String(), err)
	}

	var ab ability.A
	if err := ab.Unmarshal(val); err != nil {
		return ability.A{}, fmt.Errorf("fetch ability %s: %w", filter.ID.String(), err)
	}

	return ab, nil
}

// Delete implementation for ability in redis.
func (c *Cache) Delete(ctx context.Context, filter ability.Filter) error {
	if err := c.Do(ctx, c.B().Del().Key(key+filter.ID.String()).Build()).Error(); err != nil {
		return fmt.Errorf("remove ability %s: %w", filter.ID.String(), err)
	}

	return nil
}
