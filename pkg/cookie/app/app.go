package app

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/cookie"
	"github.com/elojah/game_03/pkg/errors"
	"github.com/gorilla/securecookie"
)

type A struct {
	memKeys []cookie.Keys

	cookie.StoreKeys

	JWT string
}

// Initial keys setup

func (a *A) Setup(ctx context.Context, n int) error {
	keys := make([]cookie.Keys, 0, n)

	if err := a.DeleteKeys(ctx, cookie.FilterKeys{
		All: true,
	}); err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		keys = append(keys, cookie.NewKeys())
	}

	if err := a.CreateKeys(ctx, keys...); err != nil {
		return err
	}

	return nil
}

// Encoding/Decoding helpers

func (a A) Encode(ctx context.Context, key string, value string) (string, error) {
	ks, err := a.ReadKeys(ctx, cookie.FilterKeys{All: true})
	if err != nil {
		return "", err
	}

	if len(ks) == 0 {
		return "", errors.ErrMissingSecureKeys{}
	}

	ck, err := securecookie.New(
		ks[len(ks)-1].Hash,
		ks[len(ks)-1].Block,
	).Encode(key, value)
	if err != nil {
		return "", err
	}

	return ck, nil
}

func (a A) Decode(ctx context.Context, key string, value string) (string, error) {
	keys, err := a.ReadKeys(ctx, cookie.FilterKeys{All: true})
	if err != nil {
		return "", err
	}

	scs := func() []securecookie.Codec {
		result := make([]securecookie.Codec, 0, len(keys))
		for _, k := range keys {
			result = append(result, securecookie.New(k.Hash, k.Block))
		}

		return result
	}()

	var s string

	if err := securecookie.DecodeMulti(key, value, &s, scs...); err != nil {
		return "", err
	}

	return s, nil
}

// Cache management

func (a *A) ReadKeys(ctx context.Context, f cookie.FilterKeys) ([]cookie.Keys, error) {
	if !f.All {
		return nil, nil
	}

	if len(a.memKeys) == 0 {
		if err := a.SyncKeys(ctx, f); err != nil {
			return nil, err
		}
	}

	return a.memKeys, nil
}

func (a *A) SyncKeys(ctx context.Context, f cookie.FilterKeys) error {
	keys, err := a.StoreKeys.ReadKeys(ctx, f)
	if err != nil {
		return err
	}

	a.memKeys = keys

	return nil
}

func (a *A) AutoSyncKeys(ctx context.Context, interval int64) error {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for range ticker.C {
		if err := a.SyncKeys(ctx, cookie.FilterKeys{
			All: true,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (a *A) JWTSecret() string {
	return a.JWT
}
