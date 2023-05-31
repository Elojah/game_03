package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filter entity.Filter

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

	if f.CellID != nil {
		clause = append(clause, `cell_id = ?`)
		args = append(args, f.CellID)
	}

	if len(f.CellIDs) > 0 {
		clause = append(clause, `cell_id IN ?`)
		args = append(args, f.CellIDs)
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

	if f.CellID != nil {
		cols = append(cols, f.CellID.String())
	}

	if f.CellIDs != nil {
		ss := make([]string, 0, len(f.CellIDs))
		for _, cid := range f.CellIDs {
			ss = append(ss, cid.String())
		}

		cols = append(cols, strings.Join(ss, "|"))
	}

	return strings.Join(cols, " - ")
}

type patch entity.Patch

func (p patch) set() (string, []any) {
	var clause []string

	var args []any

	if p.UserID != nil {
		clause = append(clause, `user_id = ?`)
		args = append(args, p.UserID)
	}

	if p.CellID != nil {
		clause = append(clause, `cell_id = ?`)
		args = append(args, p.CellID)
	}

	if p.FactionID != nil {
		clause = append(clause, `faction_id = ?`)
		args = append(args, p.FactionID)
	}

	if p.Name != nil {
		clause = append(clause, `name = ?`)
		args = append(args, *p.Name)
	}

	if p.X != nil {
		clause = append(clause, `x = ?`)
		args = append(args, *p.X)
	}

	if p.Y != nil {
		clause = append(clause, `y = ?`)
		args = append(args, *p.Y)
	}

	if p.Rot != nil {
		clause = append(clause, `rot = ?`)
		args = append(args, *p.Rot)
	}

	if p.Radius != nil {
		clause = append(clause, `radius = ?`)
		args = append(args, *p.Radius)
	}

	if p.At != nil {
		clause = append(clause, `at = ?`)
		args = append(args, *p.At)
	}

	if p.AnimationID != nil {
		clause = append(clause, `animation_id = ?`)
		args = append(args, p.AnimationID)
	}

	if p.AnimationAt != nil {
		clause = append(clause, `animation_at = ?`)
		args = append(args, *p.AnimationAt)
	}

	if p.Objects != nil {
		clause = append(clause, `objects = ?`)
		args = append(args, p.Objects)
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

func (s Store) Insert(ctx context.Context, e entity.E) error {
	q := s.Session.Query(
		`INSERT INTO main.entity (
			id, user_id, cell_id, faction_id,
			name,
			x, y, rot, radius,
			at,
			animation_id, animation_at,
			objects
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)`,
		e.ID, e.UserID, e.CellID, e.FactionID,
		e.Name,
		e.X, e.Y, e.Rot, e.Radius,
		e.At,
		e.AnimationID, e.AnimationAt,
		e.Objects,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) Update(ctx context.Context, f entity.Filter, p entity.Patch) error {
	b := strings.Builder{}
	b.WriteString(`UPDATE main.entity`)

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

func (s Store) Fetch(ctx context.Context, f entity.Filter) (entity.E, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT
		id, user_id, cell_id, faction_id,
		name,
		x, y, rot, radius,
		at,
		animation_id, animation_at,
		objects
	FROM main.entity `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var e entity.E
	if err := q.Scan(
		&e.ID, &e.UserID, &e.CellID, &e.FactionID,
		&e.Name,
		&e.X, &e.Y, &e.Rot, &e.Radius,
		&e.At,
		&e.AnimationID, &e.AnimationAt,
		&e.Objects,
	); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.E{}, gerrors.ErrNotFound{Resource: "entity", Index: filter(f).index()}
		}

		return entity.E{}, err
	}

	return e, nil
}

func (s Store) FetchMany(ctx context.Context, f entity.Filter) ([]entity.E, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT
		id, user_id, cell_id, faction_id,
		name,
		x, y, rot, radius,
		at,
		animation_id, animation_at,
		objects
	FROM main.entity `)

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

	es := make([]entity.E, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&es[i].ID, &es[i].UserID, &es[i].CellID, &es[i].FactionID,
			&es[i].Name,
			&es[i].X, &es[i].Y, &es[i].Rot, &es[i].Radius,
			&es[i].At,
			&es[i].AnimationID, &es[i].AnimationAt,
			&es[i].Objects,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return es[:i], state, nil
}

func (s Store) Delete(ctx context.Context, f entity.Filter) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.entity `)

	clause, args := filter(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
