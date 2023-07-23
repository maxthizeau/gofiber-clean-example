package exception

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}

func PanicValidation(err error) {
	if err != nil {
		panic(ValidationError{
			Message: err.Error(),
		})
	}
}
