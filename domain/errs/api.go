package errs

import "errors"

type ErrorWithCode struct {
	Code int
	Err  error
}

func (e ErrorWithCode) Error() string {
	return e.Err.Error()
}

// Unwrap returns the wrapped error, allowing for unwrapping.
func (e ErrorWithCode) Unwrap() error {
	return e.Err
}

func NewErrorWithCode(code int, msg string) *ErrorWithCode {
	return &ErrorWithCode{
		Code: code,
		Err:  errors.New(msg),
	}
}
