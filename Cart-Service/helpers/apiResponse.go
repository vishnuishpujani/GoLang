package helpers

import "Case-Study/Cart-Service/model"

func ResponseMapper(code int, message string) model.Response {
	var response model.Response
	response = model.Response{
		Code:    code,
		Message: message,
	}
	return response

}
