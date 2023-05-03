package app

import (
	"context"

	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/game_03/pkg/event"
	"github.com/elojah/game_03/pkg/ulid"
)

const (
	bufferSize = 100
)

type App struct {
	event.Cache
	event.CacheQ

	entity entity.App
}

func (a App) Subscribe(ctx context.Context, entityID ulid.ID) error {
	ch := make(chan event.Message, bufferSize)

	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	go func() {
		for {
			select {
			case _ = <-ctx.Done():
				// log
				return
			case m := <-ch:
				var ev event.E

				if err := ev.Unmarshal([]byte(m.Message)); err != nil {
					// log
				}

			}
		}
	}()

	// Subscribe is blocking
	return a.CacheQ.Subscribe(ctx, event.FilterQ{
		EntityID: entityID,
	}, func(m event.Message) {
		ch <- m
	})
}
