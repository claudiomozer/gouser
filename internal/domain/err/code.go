package err

type Code int

const (
	ErrMissingRequiredField Code = 1
	ErrInvalidField         Code = 2
	ErrUserAlreadyExists    Code = 3
	ErrUserNotExists        Code = 4
	ErrOperationNotAllowed  Code = 5
)
