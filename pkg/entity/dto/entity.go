package dto

import "github.com/elojah/game_03/pkg/errors"

const (
	maxListEntity = 20
)

func (req ListEntityReq) Check() error {
	if len(req.IDs) > maxListEntity {
		return errors.ErrInvalidNumericalRange{
			Key:   "IDs",
			Value: len(req.IDs),
			Min:   0,
			Max:   maxListEntity,
		}
	}

	return nil
}
