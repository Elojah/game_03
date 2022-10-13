package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/elojah/game_03/pkg/cookie"
	"github.com/go-redis/redis"
)

const (
	keysKey = "cookie_keys:"
)

// CreateKeys implementation for CookieKeys in redis.
func (s *Store) CreateKeys(ctx context.Context, cks ...cookie.Keys) error {
	for _, ck := range cks {
		raw, err := ck.Marshal()
		if err != nil {
			return err
		}

		if err := s.LPush(ctx, keysKey, raw).Err(); err != nil {
			return fmt.Errorf("push cookie_keys: %w", err)
		}
	}

	return nil
}

// RotateKeys implementation for CookieKeys in redis.
func (s *Store) RotateKeys(ctx context.Context, ck cookie.Keys, size int64) error {
	raw, err := ck.Marshal()
	if err != nil {
		return err
	}

	if err := s.LPush(ctx, keysKey, raw).Err(); err != nil {
		return fmt.Errorf("push cookie_keys: %w", err)
	}

	if err := s.LTrim(ctx, keysKey, 0, size-1).Err(); err != nil {
		return fmt.Errorf("trim cookie_keys: %w", err)
	}

	return nil
}

// ReadKeys implementation for CookieKeys in redis.
func (s *Store) ReadKeys(ctx context.Context, f cookie.FilterKeys) ([]cookie.Keys, error) {
	if !f.All {
		return nil, nil
	}

	vals, err := s.LRange(ctx, keysKey, 0, -1).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("read keys %s: %w", keysKey, err)
		}
	}

	result := make([]cookie.Keys, 0, len(vals))

	for _, val := range vals {
		var ck cookie.Keys

		if err := ck.Unmarshal([]byte(val)); err != nil {
			return result, fmt.Errorf("read keys: unmarshal: %w", err)
		}

		result = append(result, ck)
	}

	return result, nil
}

// DeleteKeys implementation for CookieKeys in redis.
func (s *Store) DeleteKeys(ctx context.Context, f cookie.FilterKeys) error {
	if !f.All {
		return nil
	}

	if err := s.Del(ctx, keysKey).Err(); err != nil {
		return fmt.Errorf("delete cookie_keys: %w", err)
	}

	return nil
}
