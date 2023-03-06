package app

import (
	"context"
	"strings"

	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/user"
	"google.golang.org/grpc/metadata"
)

func (a App) AuthSession(ctx context.Context) (user.Session, error) {
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

	var ses user.Session

	if err := ses.Decrypt(a.sessionKey, []byte(token)); err != nil {
		return user.Session{}, err
	}

	return ses, nil
}

func (a App) CreateSession(ctx context.Context, ses user.Session) ([]byte, error) {
	// TODO: use cache
	if err := a.StoreSession.InsertSession(ctx, ses); err != nil {
		return nil, err
	}

	return ses.Encrypt(a.sessionKey)
}

func (a App) InsertSession(ctx context.Context, ses user.Session) error {
	// TODO: use cache
	return a.StoreSession.InsertSession(ctx, ses)
}

func (a App) FetchSession(ctx context.Context, f user.FilterSession) (user.Session, error) {
	// TODO: use cache
	return a.StoreSession.FetchSession(ctx, f)
}

func (a App) DeleteSession(ctx context.Context, f user.FilterSession) error {
	// TODO: use cache
	return a.StoreSession.DeleteSession(ctx, f)
}
