package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/go-redis/redis/v8"
)

const (
	pcConnectKey = "connect_pc_entity:"

	defaultExpire = 1 * time.Second
)

// UpsertPCConnect implementation for pcConnect in redis.
func (c *Cache) UpsertPCConnect(ctx context.Context, pcc entity.PCConnect) error {
	raw, err := pcc.Marshal()
	if err != nil {
		return err
	}

	// Upsert for default 1 hour
	if err := c.Set(ctx, pcConnectKey+pcc.ID.String(), raw, defaultExpire).Err(); err != nil {
		return fmt.Errorf("upsert pc_connect %s: %w", pcc.ID.String(), err)
	}

	return nil
}

// FetchPCConnect implementation for PCConnect in redis.
func (c *Cache) FetchPCConnect(ctx context.Context, filter entity.FilterPCConnect) (entity.PCConnect, error) {
	val, err := c.Get(ctx, pcConnectKey+filter.ID.String()).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return entity.PCConnect{}, fmt.Errorf("fetch pc_connect %s: %w", filter.ID.String(), err)
		}

		return entity.PCConnect{}, gerrors.ErrNotFound{Resource: pcConnectKey, Index: filter.ID.String()}
	}

	var pcc entity.PCConnect
	if err := pcc.Unmarshal([]byte(val)); err != nil {
		return entity.PCConnect{}, fmt.Errorf("fetch pc_connect %s: %w", filter.ID.String(), err)
	}

	return pcc, nil
}

// DeletePCConnect implementation for pcConnect in redis.
func (c *Cache) DeletePCConnect(ctx context.Context, filter entity.FilterPCConnect) error {
	if err := c.Del(ctx, pcConnectKey+filter.ID.String()).Err(); err != nil {
		return fmt.Errorf("remove pc_connect %s: %w", filter.ID.String(), err)
	}

	return nil
}
