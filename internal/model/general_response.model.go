package model

type GeneralResponse struct {
	Code    int         `json:"code" binding:"required"`
	Message string      `json:"message" binding:"required"`
	Data    interface{} `json:"data" binding:"required"`
}

/*
* Create a success response with first parameter = Data, second parameter = Message
 */

func NewSuccessResponse(params ...interface{}) GeneralResponse {
	if len(params) == 0 {
		return GeneralResponse{Code: 200}
	}

	data := params[0]
	message := "Success"

	if len(params) > 1 {
		newMessgae, ok := params[1].(string)
		if ok {
			message = newMessgae
		}
	}

	return GeneralResponse{
		Code:    200,
		Message: message,
		Data:    data,
	}
}
