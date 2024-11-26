package errors

type NotAcceptableError struct {
	*superError
	message string
}

func (instance NotAcceptableError) Error() string {
	return instance.message
}

func (instance NotAcceptableError) FriendlyMessage() string {
	return instance.message
}

func (instance NotAcceptableError) Equals(err Error) bool {
	return instance.Error() == err.Error()
}

func (instance NotAcceptableError) LogLevel() int {
	return ErrorLevel
}

func NewNotAcceptableError(message string) *NotAcceptableError {
	return &NotAcceptableError{
		superError: newSuperError(),
		message:    message,
	}
}
