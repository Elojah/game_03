package redis

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/go-redis"
)

var _ entity.Cache = (*Cache)(nil)

type Cache struct {
	redis.Service
}
