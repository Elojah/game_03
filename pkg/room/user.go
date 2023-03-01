package room

import (
	"context"

	"github.com/elojah/game_03/pkg/ulid"
)

type Role int32

const (
	UnknownRole Role = iota
	Owner
	Member
)

type FilterUser struct {
	UserID  ulid.ID
	UserIDs []ulid.ID
	RoomID  ulid.ID
	RoomIDs []ulid.ID

	State []byte
	Size  int
}

type StoreUser interface {
	InsertUser(context.Context, User) error
	FetchUser(context.Context, FilterUser) (User, error)
	FetchManyUser(context.Context, FilterUser) ([]User, []byte, error)
	DeleteUser(context.Context, FilterUser) error
}

type Users []User

func (us Users) RoomIDs() []ulid.ID {
	result := make([]ulid.ID, 0, len(us))

	for _, u := range us {
		result = append(result, u.RoomID)
	}

	return result
}
