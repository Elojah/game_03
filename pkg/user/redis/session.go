package redis

import (
	"context"
	"fmt"
	"time"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/user"
	"github.com/redis/rueidis"
)

const (
	sessionKey = "session_user:"

	cacheDuration = 300
)

// InsertSession implementation for session in redis.
func (c *Cache) InsertSession(ctx context.Context, ses user.Session) error {
	raw, err := ses.Marshal()
	if err != nil {
		return err
	}

	if err := c.Do(
		ctx,
		c.B().Set().Key(sessionKey+ses.ID.String()).Value(rueidis.BinaryString(raw)).Build(),
	).Error(); err != nil {
		return fmt.Errorf("upsert session %s: %w", ses.ID.String(), err)
	}

	return nil
}

// FetchSession implementation for session in redis.
func (c *Cache) FetchSession(ctx context.Context, filter user.FilterSession) (user.Session, error) {
	val, err := c.DoCache(ctx, c.B().Get().Key(sessionKey+filter.ID.String()).Cache(), cacheDuration*time.Second).AsBytes()
	if err != nil {
		// if !errors.Is(err, rueidis.Err) {
		// 	return user.Session{}, fmt.Errorf("fetch session %s: %w", filter.ID.String(), err)
		// }
		return user.Session{}, gerrors.ErrNotFound{Resource: sessionKey, Index: filter.ID.String()}
	}

	var ses user.Session
	if err := ses.Unmarshal(val); err != nil {
		return user.Session{}, fmt.Errorf("fetch session %s: %w", filter.ID.String(), err)
	}

	return ses, nil
}

// DeleteSession implementation for session in redis.
func (c *Cache) DeleteSession(ctx context.Context, filter user.FilterSession) error {
	if err := c.Do(ctx, c.B().Del().Key(sessionKey+filter.ID.String()).Build()).Error(); err != nil {
		return fmt.Errorf("remove session %s: %w", filter.ID.String(), err)
	}

	return nil
}
