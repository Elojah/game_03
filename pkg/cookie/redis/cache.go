package redis

import (
	"github.com/elojah/game_03/pkg/cookie"
	"github.com/elojah/go-redis"
)

var _ cookie.CacheKeys = (*Cache)(nil)

type Cache struct {
	redis.Service
}
