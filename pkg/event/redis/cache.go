package redis

import (
	"github.com/elojah/game_03/pkg/event"
	"github.com/elojah/go-redis"
)

var _ event.Cache = (*Cache)(nil)

type Cache struct {
	redis.Service
}
