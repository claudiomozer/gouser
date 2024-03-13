package err

import "fmt"

type Error struct {
	code    Code
	message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("{ key: %d - msg: %s", e.code, e.message)
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Code() Code {
	return e.code
}

func MissingRequiredField(field string) *Error {
	return &Error{
		code:    ErrMissingRequiredField,
		message: fmt.Sprintf("%s is required", field),
	}
}

func InvalidField(field string) *Error {
	return &Error{
		code:    ErrInvalidField,
		message: fmt.Sprintf("%s is invalid", field),
	}
}

func New(code Code, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}
