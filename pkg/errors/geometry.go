package errors

import "fmt"

type ErrInvalidDimension struct{}

func (e ErrInvalidDimension) Error() string {
	return fmt.Sprintf("invalid dimension")
}
