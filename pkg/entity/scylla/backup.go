package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filterBackup entity.FilterBackup

func (f filterBackup) where() (string, []interface{}) {
	var clause []string

	var args []interface{}

	if f.ID != nil {
		clause = append(clause, `id = ?`)
		args = append(args, *f.ID)
	}

	if len(f.IDs) > 0 {
		clause = append(clause, `id IN ?`)
		args = append(args, f.IDs)
	}

	if f.CellID != nil {
		clause = append(clause, `cell_id = ?`)
		args = append(args, *f.CellID)
	}

	if len(f.CellIDs) > 0 {
		clause = append(clause, `cell_id IN ?`)
		args = append(args, f.CellIDs)
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

func (f filterBackup) index() string {
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

func (s Store) InsertBackup(ctx context.Context, bu entity.Backup) error {
	q := s.Session.Query(
		`INSERT INTO main.entity (
			id, user_id, cell_id,
			name,
			x, y, rot, radius,
			at,
			animation_id, animation_at
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)`,
		bu.ID, bu.UserID, bu.CellID,
		bu.Name,
		bu.X, bu.Y, bu.Rot, bu.Radius,
		bu.At,
		bu.AnimationID, bu.AnimationAt,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchBackup(ctx context.Context, f entity.FilterBackup) (entity.Backup, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT
		id, user_id, cell_id,
		name,
		x, y, rot, radius,
		at
		animation_id, animation_at
	FROM main.entity `)

	clause, args := filterBackup(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var bu entity.Backup
	if err := q.Scan(
		&bu.ID, &bu.UserID, &bu.CellID,
		&bu.Name,
		&bu.X, &bu.Y, &bu.Rot, &bu.Radius,
		&bu.At,
		&bu.AnimationID, &bu.AnimationAt,
	); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.Backup{}, gerrors.ErrNotFound{Resource: "entity_backup", Index: filterBackup(f).index()}
		}

		return entity.Backup{}, err
	}

	return bu, nil
}

func (s Store) FetchManyBackup(ctx context.Context, f entity.FilterBackup) ([]entity.Backup, []byte, error) {
	if f.Size == 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT
		id, user_id, cell_id,
		name,
		x, y, rot, radius,
		at
		animation_id, animation_at
	FROM main.entity `)

	clause, args := filterBackup(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	bus := make([]entity.Backup, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&bus[i].ID, &bus[i].UserID, &bus[i].CellID,
			&bus[i].Name,
			&bus[i].X, &bus[i].Y, &bus[i].Rot, &bus[i].Radius,
			&bus[i].At,
			&bus[i].AnimationID, &bus[i].AnimationAt,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return bus[:i], state, nil
}

func (s Store) DeleteBackup(ctx context.Context, f entity.FilterBackup) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.entity_backup `)

	clause, args := filterBackup(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
