package exception

type ValidationException struct {
	Message string
}

func NewValidationException(message string) ValidationException {
	return ValidationException{Message: message}
}

func (exception ValidationException) Error() string {
	return exception.Message
}
