package app

import (
	"context"
	"strings"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/elojah/game_03/pkg/user"
	"google.golang.org/grpc/metadata"
)

var _ user.App = (*App)(nil)

type App struct {
	user.Store
	user.StoreSession
	user.CacheSession
}

func (a App) Auth(ctx context.Context) (user.Session, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return user.Session{}, errors.ErrMissingAuth{}
	}

	token := func() string {
		if len(md["token"]) == 0 {
			return ""
		}

		return strings.Trim(md["token"][0], " ")
	}()
	if token == "" {
		return user.Session{}, errors.ErrMissingAuth{}
	}

	id, err := ulid.Parse(token)
	if err != nil {
		return user.Session{}, errors.ErrInvalidCredentials{}
	}

	return a.StoreSession.FetchSession(ctx, user.FilterSession{ID: id})
}

func (a App) FetchSession(ctx context.Context, f user.FilterSession) (user.Session, error) {
	return a.StoreSession.FetchSession(ctx, f)
}

func (a App) DeleteSession(ctx context.Context, f user.FilterSession) error {
	return a.StoreSession.DeleteSession(ctx, f)
}
