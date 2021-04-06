package errors

import "fmt"

type ErrNullRequest struct{}

func (e ErrNullRequest) Error() string {
	return fmt.Sprintf("null request")
}

type ErrMissingAuth struct{}

func (e ErrMissingAuth) Error() string {
	return fmt.Sprintf("missing authentication")
}

type ErrInvalidCredentials struct{}

func (e ErrInvalidCredentials) Error() string {
	return fmt.Sprintf("invalid credentials")
}
