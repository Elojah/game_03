package redis

import (
	"github.com/elojah/game_03/pkg/user"
	"github.com/elojah/go-redis"
)

var _ user.CacheSession = (*Cache)(nil)

type Cache struct {
	redis.Service
}
