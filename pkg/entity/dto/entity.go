package dto

import "github.com/elojah/game_03/pkg/errors"

const (
	maxEntityIDs = 20
	maxCellIDs   = 9
)

func (req ListEntityReq) Check() error {
	if len(req.IDs) > maxEntityIDs {
		return errors.ErrInvalidNumericalRange{
			Key:   "IDs",
			Value: len(req.IDs),
			Min:   0,
			Max:   maxEntityIDs,
		}
	}

	if len(req.CellIDs) > maxCellIDs {
		return errors.ErrInvalidNumericalRange{
			Key:   "CellIDs",
			Value: len(req.CellIDs),
			Min:   0,
			Max:   maxCellIDs,
		}
	}

	return nil
}
