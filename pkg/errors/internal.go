package errors

import "fmt"

type ErrNotImplemented struct {
	Version string
}

func (err ErrNotImplemented) Error() string {
	return fmt.Sprintf("not implemented yet, planned for %s", err.Version)
}

type ErrMissingRequiredFilter struct {
	Filter string
}

func (err ErrMissingRequiredFilter) Error() string {
	return fmt.Sprintf("filter %s must be set", err.Filter)
}
