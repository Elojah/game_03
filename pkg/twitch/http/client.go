package http

import (
	"net/http"

	"github.com/elojah/game_03/pkg/twitch"
)

var _ twitch.Client = (*Client)(nil)

type Client struct {
	*http.Client
}
