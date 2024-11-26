package errors

var _ Error = (*UnauthorizedError)(nil)

type UnauthorizedError struct {
	*superError
	message string
	err     error
}

func (instance UnauthorizedError) FriendlyMessage() string {
	return instance.message
}

func (instance UnauthorizedError) Equals(err Error) bool {
	return instance.Error() == err.Error()
}

func (instance UnauthorizedError) Error() string {
	return instance.err.Error()
}

func (instance UnauthorizedError) LogLevel() int {
	return ErrorLevel
}

func NewUnauthorizedError(msg string, err error) *UnauthorizedError {
	return &UnauthorizedError{newSuperError(), msg, err}
}
