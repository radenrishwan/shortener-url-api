package exception

type IsExistException struct {
	Message string
}

func NewIsExistException(message string) IsExistException {
	return IsExistException{Message: message}
}

func (exception IsExistException) Error() string {
	return exception.Message
}
