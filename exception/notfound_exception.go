package exception

type NotFoundException struct {
	Message string
}

func NewNotFoundException(message string) NotFoundException {
	return NotFoundException{Message: message}
}

func (exception NotFoundException) Error() string {
	return exception.Message
}
