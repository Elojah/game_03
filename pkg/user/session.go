package user

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
	"github.com/xxtea/xxtea-go/xxtea"
)

// sessionKey is a weak security key but fast and constant.
var sessionKey = []byte("01GG4HA3DM43KT1F8YRGZ5EZZ5")

type FilterSession struct {
	ID ulid.ID
}

type CacheSession interface {
	UpsertSession(context.Context, Session) error
	FetchSession(context.Context, FilterSession) (Session, error)
	DeleteSession(context.Context, FilterSession) error
}

type StoreSession interface {
	InsertSession(context.Context, Session) error
	FetchSession(context.Context, FilterSession) (Session, error)
	DeleteSession(context.Context, FilterSession) error
}

func (ses Session) Encrypt() ([]byte, error) {
	raw, err := ses.Marshal()
	if err != nil {
		return nil, nil
	}

	return xxtea.Encrypt(raw, sessionKey), nil
}

func (ses *Session) Decrypt(s []byte) error {
	raw := xxtea.Decrypt(s, sessionKey)

	return ses.Unmarshal(raw)
}
