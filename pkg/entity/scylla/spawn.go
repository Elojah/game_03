package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filterSpawn entity.FilterSpawn

func (f filterSpawn) where() (string, []any) {
	var clause []string

	var args []any

	if f.EntityID != nil {
		clause = append(clause, `entity_id = ?`)
		args = append(args, f.EntityID)
	}

	if len(f.EntityIDs) > 0 {
		clause = append(clause, `entity_id IN ?`)
		args = append(args, f.EntityIDs)
	}

	if f.SpawnID != nil {
		clause = append(clause, `spawn_id = ?`)
		args = append(args, f.SpawnID)
	}

	if len(f.SpawnIDs) > 0 {
		clause = append(clause, `spawn_id IN ?`)
		args = append(args, f.SpawnIDs)
	}

	b := strings.Builder{}

	if f.All {
		return b.String(), []any{}
	}

	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	return b.String(), args
}

func (f filterSpawn) index() string {
	var cols []string

	if f.EntityID != nil {
		cols = append(cols, f.EntityID.String())
	}

	if f.EntityIDs != nil {
		ss := make([]string, 0, len(f.EntityIDs))
		for _, id := range f.EntityIDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.SpawnID != nil {
		cols = append(cols, f.SpawnID.String())
	}

	if f.SpawnIDs != nil {
		ss := make([]string, 0, len(f.SpawnIDs))
		for _, id := range f.SpawnIDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.All {
		cols = append(cols, "all")
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertSpawn(ctx context.Context, t entity.Spawn) error {
	q := s.Session.Query(
		`INSERT INTO main.entity_spawn (entity_id, spawn_id) VALUES (?, ?)`,
		t.EntityID,
		t.SpawnID,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchSpawn(ctx context.Context, f entity.FilterSpawn) (entity.Spawn, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT entity_id, spawn_id FROM main.entity_spawn `)

	clause, args := filterSpawn(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var t entity.Spawn
	if err := q.Scan(&t.EntityID, &t.SpawnID); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.Spawn{}, gerrors.ErrNotFound{Resource: "entity_spawn", Index: filterSpawn(f).index()}
		}

		return entity.Spawn{}, err
	}

	return t, nil
}

func (s Store) FetchManySpawn(ctx context.Context, f entity.FilterSpawn) ([]entity.Spawn, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT entity_id, spawn_id FROM main.entity_spawn `)

	clause, args := filterSpawn(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	ts := make([]entity.Spawn, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&ts[i].EntityID,
			&ts[i].SpawnID,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return ts[:i], state, nil
}

func (s Store) DeleteSpawn(ctx context.Context, f entity.FilterSpawn) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.entity_spawn `)

	clause, args := filterSpawn(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
