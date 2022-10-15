package app

import (
	"context"
	"strings"
	"time"

	"github.com/elojah/game_03/pkg/cookie"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/ulid"
	"github.com/elojah/game_03/pkg/user"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/metadata"
)

var _ user.App = (*App)(nil)

type App struct {
	user.Store
	user.StoreSession
	user.CacheSession

	Cookie cookie.App

	secret string
}

func (a App) CreateJWT(ctx context.Context, u user.U) (string, error) {
	// #Create JWT
	id := ulid.NewID()

	// use cookie rotation encoding to generate a rotating secret for JWT
	secret, err := a.Cookie.Encode(ctx, "jwt_secret", id.String())
	if err != nil {
		return "", err
	}

	now := time.Now()
	claims := &user.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "spc",
			Subject:   u.ID.String(),
			Audience:  jwt.ClaimStrings{"log"},
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        secret,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(id.String()))
	if err != nil {
		return "", err
	}

	ck, err := a.Cookie.Encode(ctx, "token", ss)
	if err != nil {
		return "", err
	}

	return ck, nil
}

func (a App) ReadJWT(ctx context.Context, token string) (user.U, error) {
	ck, err := a.Cookie.Decode(ctx, "token", token)
	if err != nil {
		return user.U{}, err
	}

	t, err := jwt.ParseWithClaims(ck, &user.Claims{}, func(t *jwt.Token) (interface{}, error) {
		claims, ok := t.Claims.(*user.Claims)
		if !ok {
			return user.U{}, errors.ErrInvalidClaims{}
		}

		secret, err := a.Cookie.Decode(ctx, "jwt_secret", claims.RegisteredClaims.ID)
		if err != nil {
			return "", err
		}

		return []byte(secret), nil
	})
	if err != nil {
		return user.U{}, err
	}

	// check token validity
	claims, ok := t.Claims.(*user.Claims)
	if !ok {
		return user.U{}, errors.ErrInvalidClaims{}
	}

	if err := claims.Valid(); err != nil {
		return user.U{}, errors.ErrInvalidClaims{Err: err}
	}

	id, err := ulid.Parse(claims.Subject)
	if err != nil {
		return user.U{}, errors.ErrInvalidClaims{Err: err}
	}

	return user.U{
		ID: id,
	}, nil
}

func (a App) Auth(ctx context.Context) (user.U, error) {
	// read & parse token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return user.U{}, errors.ErrMissingAuth{}
	}

	token := func() string {
		if len(md["token"]) == 0 {
			return ""
		}

		return strings.Trim(md["token"][0], " ")
	}()

	if token == "" {
		return user.U{}, errors.ErrMissingAuth{}
	}

	return a.ReadJWT(ctx, token)
}

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
