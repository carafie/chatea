package errorsx

import (
	"errors"
)

type Error struct {
	Code    string
	Message string
}

func New(code, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}

func From(err error) (Error, bool) {
	if err, ok := errors.AsType[Error](err); ok {
		return err, true
	}
	return Error{}, false
}

func (e Error) Error() string {
	return e.Message
}
