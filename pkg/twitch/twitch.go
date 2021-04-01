package twitch

import (
	"context"

	"golang.org/x/oauth2"
)

type Auth struct {
	Token    string
	ClientID string
}

type Client interface {
	// Token
	GetToken(context.Context, string, oauth2.Config) (Token, error)

	// User
	GetUsers(context.Context, Auth, UserFilter, func(User) error) error
}

type App interface {
	Client
	OAuth() oauth2.Config
}
