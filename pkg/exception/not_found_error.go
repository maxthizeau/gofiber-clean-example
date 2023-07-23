package exception

type NotFoundError struct {
	Message string
}

func (notFoundError NotFoundError) Error() string {
	return notFoundError.Message
}

func PanicNotFound(err error) {
	if err != nil {
		panic(NotFoundError{
			Message: err.Error(),
		})
	}
}
