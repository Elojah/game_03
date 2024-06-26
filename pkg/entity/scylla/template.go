package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filterTemplate entity.FilterTemplate

func (f filterTemplate) where() (string, []any) {
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

	if f.Name != nil {
		clause = append(clause, `name = ?`)
		args = append(args, f.Name)
	}

	if len(f.Names) > 0 {
		clause = append(clause, `name IN ?`)
		args = append(args, f.Names)
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

func (f filterTemplate) index() string {
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

	if f.Name != nil {
		cols = append(cols, *f.Name)
	}

	if f.Names != nil {
		ss := make([]string, 0, len(f.Names))
		for _, name := range f.Names {
			ss = append(ss, name)
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.All {
		cols = append(cols, "all")
	}

	return strings.Join(cols, " - ")
}

func (s Store) InsertTemplate(ctx context.Context, t entity.Template) error {
	q := s.Session.Query(
		`INSERT INTO main.entity_template (entity_id, name) VALUES (?, ?)`,
		t.EntityID,
		t.Name,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchTemplate(ctx context.Context, f entity.FilterTemplate) (entity.Template, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT entity_id, name FROM main.entity_template `)

	clause, args := filterTemplate(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var t entity.Template
	if err := q.Scan(&t.EntityID, &t.Name); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.Template{}, gerrors.ErrNotFound{Resource: "template", Index: filterTemplate(f).index()}
		}

		return entity.Template{}, err
	}

	return t, nil
}

func (s Store) FetchManyTemplate(ctx context.Context, f entity.FilterTemplate) ([]entity.Template, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT entity_id, name FROM main.entity_template `)

	clause, args := filterTemplate(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	ts := make([]entity.Template, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&ts[i].EntityID,
			&ts[i].Name,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return ts[:i], state, nil
}

func (s Store) DeleteTemplate(ctx context.Context, f entity.FilterTemplate) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.entity_template `)

	clause, args := filterTemplate(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
