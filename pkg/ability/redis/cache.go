package redis

import (
	"github.com/elojah/game_03/pkg/ability"
	"github.com/elojah/go-redis"
)

var _ ability.Cache = (*Cache)(nil)

type Cache struct {
	redis.Service
}
