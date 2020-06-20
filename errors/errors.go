package errors

import (
	"fmt"
)

type Error struct {
	Err        error  `json:"err"`
	StatusCode int    `json:"status_code"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("err: %s, statusCode: %d, code: %d, message: %s", e.Err, e.StatusCode, e.Code, e.Message)
}

func New(err error, statusCode int, code int, msg string) *Error {
	return &Error{
		Err:        err,
		StatusCode: statusCode,
		Code:       code,
		Message:    msg,
	}
}
