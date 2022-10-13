package app

import (
	"context"
	"time"

	"github.com/elojah/game_03/pkg/cookie"
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
