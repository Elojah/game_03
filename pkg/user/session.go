package user

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
	"github.com/xxtea/xxtea-go/xxtea"
)

type FilterSession struct {
	ID ulid.ID
}

type CacheSession interface {
	InsertSession(context.Context, Session) error
	FetchSession(context.Context, FilterSession) (Session, error)
	DeleteSession(context.Context, FilterSession) error
}

type StoreSession interface {
	InsertSession(context.Context, Session) error
	FetchSession(context.Context, FilterSession) (Session, error)
	DeleteSession(context.Context, FilterSession) error
}

func (ses Session) Encrypt(key []byte) ([]byte, error) {
	raw, err := ses.Marshal()
	if err != nil {
		return nil, nil
	}

	return xxtea.Encrypt(raw, key), nil
}

func (ses *Session) Decrypt(key []byte, s []byte) error {
	raw := xxtea.Decrypt(s, key)

	return ses.Unmarshal(raw)
}
