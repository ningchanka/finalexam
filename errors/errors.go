package errors

import "fmt"

type Error struct {
	Err     error
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("err: %s, code: %d, message: %s", e.Err, e.Code, e.Message)
}

func New(err error, code int, msg string) *Error {
	return &Error{
		Err:     err,
		Code:    code,
		Message: msg,
	}
}
