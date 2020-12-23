package redis

import (
	"github.com/elojah/game_03/pkg/entity"
	"github.com/elojah/go-redis"
)

var _ entity.Store = (*Store)(nil)

type Store struct {
	redis.Service
}
