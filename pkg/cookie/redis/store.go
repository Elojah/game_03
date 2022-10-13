package redis

import (
	"github.com/elojah/game_03/pkg/cookie"
	"github.com/elojah/go-redis"
)

var _ cookie.StoreKeys = (*Store)(nil)

type Store struct {
	redis.Service
}
