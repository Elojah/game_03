package redis

import (
	"context"
	"fmt"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/redis/rueidis"
)

const (
	abilityKey = "entity_ability:"
)

// InsertCacheAbility implementation for ability in redis.
func (c *Cache) InsertCacheAbility(ctx context.Context, eab entity.Ability) error {
	raw, err := eab.Marshal()
	if err != nil {
		return err
	}

	if err := c.Do(
		ctx,
		c.B().Set().Key(abilityKey+eab.AbilityID.String()).Value(rueidis.BinaryString(raw)).Build(),
	).Error(); err != nil {
		return fmt.Errorf("upsert ability %s: %w", eab.AbilityID.String(), err)
	}

	return nil
}

// FetchCacheAbility implementation for ability in redis.
func (c *Cache) FetchCacheAbility(ctx context.Context, filter entity.FilterAbility) (entity.Ability, error) {
	val, err := c.DoCache(ctx, c.B().Get().Key(abilityKey+filter.AbilityID.String()).Cache(), 0).AsBytes()
	if err != nil {
		return entity.Ability{}, fmt.Errorf("fetch ability %s: %w", filter.AbilityID.String(), err)
	}

	var eab entity.Ability
	if err := eab.Unmarshal(val); err != nil {
		return entity.Ability{}, fmt.Errorf("fetch ability %s: %w", filter.AbilityID.String(), err)
	}

	return eab, nil
}

// DeleteCacheAbility implementation for ability in redis.
func (c *Cache) DeleteCacheAbility(ctx context.Context, filter entity.FilterAbility) error {
	if err := c.Do(ctx, c.B().Del().Key(abilityKey+filter.AbilityID.String()).Build()).Error(); err != nil {
		return fmt.Errorf("remove ability %s: %w", filter.AbilityID.String(), err)
	}

	return nil
}
