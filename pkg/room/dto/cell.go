package dto

import "github.com/elojah/game_03/pkg/errors"

const (
	maxListCell = 9
)

func (req ListCellReq) Check() error {
	if len(req.IDs) > maxListCell {
		return errors.ErrInvalidNumericalRange{
			Key:   "IDs",
			Value: len(req.IDs),
			Min:   0,
			Max:   maxListCell,
		}
	}

	return nil
}
