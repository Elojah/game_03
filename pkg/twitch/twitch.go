package twitch

import (
	"context"
)

type Auth struct {
	Token    string
	ClientID string
}

type Client interface {
	GetUsers(context.Context, Auth, UserFilter, func(User) error) error
}

type App interface {
	Client
	ClientID() string
}
