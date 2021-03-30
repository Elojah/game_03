package scylla

import (
	"context"
	"strings"

	"github.com/elojah/game_03/pkg/user"
)

type userScylla user.U

func (u userScylla) cols() string {
	return `
	id,
	twitch_id
	`
}

func (u userScylla) args() []interface{} {
	return []interface{}{
		u.ID,
		u.TwitchID,
	}
}

type filter user.Filter

func (f filter) where() (string, []interface{}) {
	var clause []string

	var args []interface{}

	if f.ID != nil {
		clause = append(clause, " id = ? ")
	}

	if f.TwitchID != nil {
		clause = append(clause, " twitch_id = ? ")
	}

	b := strings.Builder{}

	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	b.WriteString(";")

	return b.String(), args

}

func Insert(ctx context.Context, u user.U) error {
	return nil
}

func Fetch(ctx context.Context, f user.Filter) (user.U, error) {
	return user.U{}, nil
}
