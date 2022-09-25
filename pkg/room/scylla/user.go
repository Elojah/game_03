package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/room"
	"github.com/gocql/gocql"
)

type filterUser room.FilterUser

func (f filterUser) where() (string, []interface{}) {
	var clause []string

	var args []interface{}

	if f.UserID != nil {
		clause = append(clause, `user_id = ?`)
		args = append(args, f.UserID)
	}

	if len(f.UserIDs) > 0 {
		clause = append(clause, `user_id IN ?`)
		args = append(args, f.UserIDs)
	}

	if f.RoomID != nil {
		clause = append(clause, `room_id = ?`)
		args = append(args, f.RoomID)
	}

	if len(f.RoomIDs) > 0 {
		clause = append(clause, `room_id IN ?`)
		args = append(args, f.RoomIDs)
	}

	b := strings.Builder{}
	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	return b.String(), args
}

func (f filterUser) index() string {
	var cols []string

	if f.UserID != nil {
		cols = append(cols, f.UserID.String())
	}

	if f.UserIDs != nil {
		ss := make([]string, 0, len(f.UserIDs))
		for _, id := range f.UserIDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.RoomID != nil {
		cols = append(cols, f.RoomID.String())
	}

	if f.RoomIDs != nil {
		ss := make([]string, 0, len(f.RoomIDs))
		for _, id := range f.RoomIDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertUser(ctx context.Context, c room.User) error {
	q := s.Session.Query(
		`INSERT INTO main.room_user (user_id, room_id, role) VALUES (?, ?, ?)`,
		c.UserID,
		c.RoomID,
		c.Role,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchUser(ctx context.Context, f room.FilterUser) (room.User, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT user_id, room_id, role FROM main.room_user `)

	clause, args := filterUser(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var c room.User
	if err := q.Scan(&c.UserID, &c.RoomID, &c.Role); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return room.User{}, gerrors.ErrNotFound{Resource: "room_user", Index: filterUser(f).index()}
		}

		return room.User{}, err
	}

	return c, nil
}

func (s Store) FetchManyUser(ctx context.Context, f room.FilterUser) ([]room.User, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT user_id, room_id, role FROM main.room_user `)

	clause, args := filterUser(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	cells := make([]room.User, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&cells[i].UserID,
			&cells[i].RoomID,
			&cells[i].Role,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return cells[:i], state, nil
}

func (s Store) DeleteUser(ctx context.Context, f room.FilterUser) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.room_user `)

	clause, args := filterUser(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
