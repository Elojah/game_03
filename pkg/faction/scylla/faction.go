package scylla

import (
	"context"
	"errors"
	"strings"

	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/elojah/game_03/pkg/faction"
	"github.com/gocql/gocql"
)

type filter faction.Filter

func (f filter) where() (string, []any) {
	var clause []string

	var args []any

	var allowFiltering bool

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, f.ID)
	}

	if len(f.IDs) > 0 {
		clause = append(clause, `id IN ?`)
		args = append(args, f.IDs)
	}

	if f.WorldID != nil {
		clause = append(clause, `world_id = ?`)
		args = append(args, f.WorldID)
	}

	if len(f.WorldIDs) > 0 {
		clause = append(clause, `world_id IN ?`)
		args = append(args, f.WorldIDs)
		allowFiltering = true
	}

	b := strings.Builder{}
	b.WriteString(" WHERE ")

	if len(clause) == 0 {
		b.WriteString("false")
	} else {
		b.WriteString(strings.Join(clause, " AND "))
	}

	if allowFiltering {
		b.WriteString(" ALLOW FILTERING ")
	}

	return b.String(), args
}

func (f filter) index() string {
	var cols []string

	if f.ID != nil {
		cols = append(cols, f.ID.String())
	}

	if f.IDs != nil {
		ss := make([]string, 0, len(f.IDs))
		for _, id := range f.IDs {
			ss = append(ss, id.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	if f.WorldID != nil {
		cols = append(cols, f.WorldID.String())
	}

	if f.WorldIDs != nil {
		ss := make([]string, 0, len(f.WorldIDs))
		for _, cid := range f.WorldIDs {
			ss = append(ss, cid.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

type patch faction.Patch

func (p patch) set() (string, []any) {
	var clause []string

	var args []any

	if p.Name != nil {
		clause = append(clause, `name = ?`)
		args = append(args, *p.Name)
	}

	if p.Icon != nil {
		clause = append(clause, `icon = ?`)
		args = append(args, *p.Icon)
	}

	b := strings.Builder{}
	b.WriteString(" SET ")

	if len(clause) == 0 {
		b.WriteString("0") // raise error update
	} else {
		b.WriteString(strings.Join(clause, ", "))
	}

	return b.String(), args
}

func (s Store) Insert(ctx context.Context, fac faction.F) error {
	q := s.Session.Query(
		`INSERT INTO main.faction (
			id, world_id,
			name, icon, max
		) VALUES (
			?, ?, ?, ?, ?
		)`,
		fac.ID, fac.WorldID,
		fac.Name, fac.Icon, fac.Max,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) Update(ctx context.Context, f faction.Filter, p faction.Patch) error {
	b := strings.Builder{}
	b.WriteString(`UPDATE main.faction`)

	set, sArgs := patch(p).set()
	b.WriteString(set)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), append(sArgs, args...)...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) Fetch(ctx context.Context, f faction.Filter) (faction.F, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT
		id, world_id,
		name, icon, max
	FROM main.faction `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var fac faction.F
	if err := q.Scan(
		&fac.ID, &fac.WorldID,
		&fac.Name, &fac.Icon, &fac.Max,
	); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return faction.F{}, gerrors.ErrNotFound{Resource: "faction", Index: filter(f).index()}
		}

		return faction.F{}, err
	}

	return fac, nil
}

func (s Store) FetchMany(ctx context.Context, f faction.Filter) ([]faction.F, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT
		id, world_id,
		name, icon, max
	FROM main.faction `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	facs := make([]faction.F, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&facs[i].ID, &facs[i].WorldID,
			&facs[i].Name, &facs[i].Icon, &facs[i].Max,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return facs[:i], state, nil
}

func (s Store) Delete(ctx context.Context, f faction.Filter) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.faction `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
