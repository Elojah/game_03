package dto

import "github.com/elojah/game_03/pkg/errors"

// random assumption, no technical limitation evaluated.
const (
	minWorldIDs = 1
	maxWorldIDs = 3
)

func (req ListWorldReq) Check() error {
	if len(req.IDs) > maxWorldIDs {
		return errors.ErrInvalidNumericalRange{
			Key:   "IDs length",
			Value: len(req.IDs),

			Min: minWorldIDs,
			Max: maxWorldIDs,
		}
	}

	return nil
}
