package errors

import "fmt"

type ErrInvalidStat struct {
	Value string
}

func (e ErrInvalidStat) Error() string {
	return fmt.Sprintf("invalid stat %s", e.Value)
}

type ErrInvalidOperator struct {
	Value string
}

func (e ErrInvalidOperator) Error() string {
	return fmt.Sprintf("invalid operator %s", e.Value)
}
