package scylla

import (
	"context"
	"errors"
	"strings"

	"github.com/elojah/game_03/pkg/entity"
	gerrors "github.com/elojah/game_03/pkg/errors"
	"github.com/gocql/gocql"
)

type filterAnimation entity.FilterAnimation

func (f filterAnimation) where() (string, []any) {
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

	if f.EntityID != nil {
		clause = append(clause, `entity_id = ?`)
		args = append(args, f.EntityID)
	}

	if len(f.EntityIDs) > 0 {
		clause = append(clause, `entity_id IN ?`)
		args = append(args, f.EntityIDs)
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

func (f filterAnimation) index() string {
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

	return strings.Join(cols, " - ")
}

func (s Store) InsertAnimation(ctx context.Context, an entity.Animation) error {
	q := s.Session.Query(
		`INSERT INTO main.entity_animation (
			id, entity_id, sheet_id, duplicate_id,
			name,
			start, end, sequence, rate,
			frame_width, frame_height, frame_start, frame_end, frame_margin, frame_spacing,
			repeat, delay, duration, show_and_hide
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		an.ID, an.EntityID, an.SheetID, an.DuplicateID,
		an.Name,
		an.Start, an.End, an.Sequence, an.Rate,
		an.FrameWidth, an.FrameHeight, an.FrameStart, an.FrameEnd, an.FrameMargin, an.FrameSpacing,
		an.Repeat, an.Delay, an.Duration, an.ShowAndHide,
	).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (s Store) FetchAnimation(ctx context.Context, f entity.FilterAnimation) (entity.Animation, error) {
	b := strings.Builder{}
	b.WriteString(`SELECT
		id, entity_id, sheet_id, duplicate_id,
		name,
		start, end, sequence, rate,
		frame_width, frame_height, frame_start, frame_end, frame_margin, frame_spacing,
		repeat, delay, duration, show_and_hide
	FROM main.entity_animation `)

	clause, args := filterAnimation(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	var an entity.Animation
	if err := q.Scan(
		&an.ID, &an.EntityID, &an.SheetID, &an.DuplicateID,
		&an.Name,
		&an.Start, &an.End, &an.Sequence, &an.Rate,
		&an.FrameWidth, &an.FrameHeight, &an.FrameStart, &an.FrameEnd, &an.FrameMargin, &an.FrameSpacing,
		&an.Repeat, &an.Delay, &an.Duration, &an.ShowAndHide,
	); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return entity.Animation{}, gerrors.ErrNotFound{Resource: "animation", Index: filterAnimation(f).index()}
		}

		return entity.Animation{}, err
	}

	return an, nil
}

func (s Store) FetchManyAnimation(ctx context.Context, f entity.FilterAnimation) ([]entity.Animation, []byte, error) {
	if f.Size <= 0 {
		return nil, nil, nil
	}

	b := strings.Builder{}
	b.WriteString(`SELECT
		id, entity_id, sheet_id, duplicate_id,
		name,
		start, end, sequence, rate,
		frame_width, frame_height, frame_start, frame_end, frame_margin, frame_spacing,
		repeat, delay, duration, show_and_hide
	FROM main.entity_animation `)

	clause, args := filterAnimation(f).where()
	b.WriteString(clause)

	iter := s.Session.Query(b.String(), args...).
		WithContext(ctx).
		PageState(f.State).
		PageSize(f.Size).
		Iter()

	defer iter.Close()

	state := iter.PageState()

	scanner := iter.Scanner()

	ans := make([]entity.Animation, f.Size)

	var i int

	for ; scanner.Next(); i++ {
		if err := scanner.Scan(
			&ans[i].ID, &ans[i].EntityID, &ans[i].SheetID, &ans[i].DuplicateID,
			&ans[i].Name,
			&ans[i].Start, &ans[i].End, &ans[i].Sequence, &ans[i].Rate,
			&ans[i].FrameWidth, &ans[i].FrameHeight,
			&ans[i].FrameStart, &ans[i].FrameEnd,
			&ans[i].FrameMargin, &ans[i].FrameSpacing,
			&ans[i].Repeat, &ans[i].Delay, &ans[i].Duration, &ans[i].ShowAndHide,
		); err != nil {
			return nil, nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return ans[:i], state, nil
}

func (s Store) DeleteAnimation(ctx context.Context, f entity.FilterAnimation) error {
	b := strings.Builder{}
	b.WriteString(`DELETE FROM main.entity_animation `)

	clause, args := filterAnimation(f).where()
	b.WriteString(clause)

	q := s.Session.Query(b.String(), args...).WithContext(ctx)

	defer q.Release()

	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}
