package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/user"
	"github.com/go-redis/redis/v8"
)

const sessionKey = "session_user:"

const sessionTTL = 2 * time.Hour

// UpsertSession implementation for session in redis.
func (c *Cache) UpsertSession(ctx context.Context, ses user.Session) error {
	raw, err := ses.Marshal()
	if err != nil {
		return err
	}

	if err := c.Set(ctx, sessionKey+ses.ID.String(), raw, sessionTTL).Err(); err != nil {
		return fmt.Errorf("upsert session %s: %w", ses.ID.String(), err)
	}

	return nil
}

// FetchSession implementation for session in redis.
func (c *Cache) FetchSession(ctx context.Context, filter user.FilterSession) (user.Session, error) {
	val, err := c.Get(ctx, sessionKey+filter.ID.String()).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return user.Session{}, fmt.Errorf("fetch session %s: %w", filter.ID.String(), err)
		}

		return user.Session{}, gerrors.ErrNotFound{Resource: sessionKey, Index: filter.ID.String()}
	}

	var ses user.Session
	if err := ses.Unmarshal([]byte(val)); err != nil {
		return user.Session{}, fmt.Errorf("fetch session %s: %w", filter.ID.String(), err)
	}

	return ses, nil
}

// DeleteSession implementation for session in redis.
func (c *Cache) DeleteSession(ctx context.Context, filter user.FilterSession) error {
	if err := c.Del(ctx, sessionKey+filter.ID.String()).Err(); err != nil {
		return fmt.Errorf("remove session %s: %w", filter.ID.String(), err)
	}

	return nil
}
