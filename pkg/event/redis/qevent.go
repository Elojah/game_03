package redis

import (
	"context"
	"fmt"

	"github.com/elojah/game_03/pkg/event"
	"github.com/redis/rueidis"
)

const (
	qeventKey = "qevent:"
)

func (c Cache) Publish(ctx context.Context, ev event.E) error {
	raw, err := ev.Marshal()
	if err != nil {
		return err
	}

	if err := c.Do(
		ctx,
		c.B().Publish().Channel(qeventKey+ev.EntityID.String()).Message(rueidis.BinaryString(raw)).Build(),
	).Error(); err != nil {
		return fmt.Errorf("publish event %s to %s: %w", ev.ID.String(), ev.EntityID.String(), err)
	}

	return nil
}

func (c Cache) Subscribe(ctx context.Context, f event.FilterQ, callback func(event.Message)) error {
	if err := c.Receive(
		ctx,
		c.B().Subscribe().Channel(qeventKey+f.EntityID.String()).Build(),
		callback,
	); err != nil {
		return fmt.Errorf("receive event on %s: %w", f.EntityID.String(), err)
	}

	return nil
}
