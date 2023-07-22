package exception

type UnauthorizedError struct {
	Message string
}

func (unauthorizedError UnauthorizedError) Error() string {
	return unauthorizedError.Message
}

func PanicUnauthorized(err error) {
	if err != nil {
		panic(UnauthorizedError{
			Message: err.Error(),
		})
	}
}
