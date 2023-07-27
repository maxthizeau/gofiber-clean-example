package exception

type BadRequestError struct {
	Message string
}

func (badRequestError BadRequestError) Error() string {
	return badRequestError.Message
}

func PanicBadRequest(err error) {
	if err != nil {
		panic(BadRequestError{
			Message: err.Error(),
		})
	}
}
