package model

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

type errNotFound struct{}

func (errNotFound) Error() string { return "not found" }

type errBadRequest struct{}

func (errBadRequest) Error() string { return "bad request" }

type errUnauthorized struct{}

func (errUnauthorized) Error() string { return "not found" }

type errForbidden struct{}

func (errForbidden) Error() string { return "bad request" }

func NotFound(msg string) error {
	return fmt.Errorf("%w: %s", ErrNotFound, msg)
}

func BadRequest(msg string) error {
	return fmt.Errorf("%w: %s", ErrBadRequest, msg)
}

func Unauthorized(msg string) error {
	return fmt.Errorf("%w: %s", ErrUnauthorized, msg)
}

func Forbidden(msg string) error {
	return fmt.Errorf("%w: %s", ErrForbidden, msg)
}
