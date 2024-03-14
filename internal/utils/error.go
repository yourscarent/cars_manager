package utils

import "errors"

type CustomError struct {
	Status int
	Msg    string
}

func (c CustomError) Error() string {
	return c.Msg
}

const (
	ErrInternal int = iota
	ErrNotFound
	ErrBadRequest

	InternalError = "internal error"
)

func Error(status int, msg string) error {
	return &CustomError{
		Status: status,
		Msg:    msg,
	}
}

// FromError second return argument means
// if true -> internal or unknown error
// else -> user error
func FromError(err error) (string, bool) {
	var e *CustomError
	if !errors.As(err, &e) {
		return err.Error(), true
	}

	switch e.Status {
	case ErrBadRequest:
		return e.Msg, false
	case ErrNotFound:
		return e.Msg, false
	default:
		return err.Error(), true
	}
}
