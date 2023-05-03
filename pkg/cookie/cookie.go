package cookie

import "context"

const (
	keyLength = 64
	NKeys     = 5
)

// App

type App interface {
	StoreKeys

	Encode(context.Context, string, string) (string, error)
	Decode(context.Context, string, string) (string, error)
}
