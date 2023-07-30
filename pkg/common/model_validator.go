package common

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
)

type ValidationResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Validate(modelInterface interface{}) {
	validate := validator.New()

	err := validate.Struct(modelInterface)

	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Tag(),
			})
		}

		jsonMessage, err := json.Marshal(messages)
		exception.PanicLogging(err)

		panic(exception.ValidationError{
			Message: string(jsonMessage),
		})
	}
}
