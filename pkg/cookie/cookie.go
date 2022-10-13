package cookie

import "context"

const (
	keyLength = 64
	NKeys     = 5
)

// App

type App interface {
	StoreKeys

	ReadKeys(context.Context, FilterKeys) ([]Keys, error)
	SyncKeys(context.Context, FilterKeys) error
	AutoSyncKeys(context.Context, int64) error

	JWTSecret() string
}
