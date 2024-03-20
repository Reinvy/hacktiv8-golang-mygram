package helper

import "mygram/model"

func CreateResponse(isSuccess bool, data any, errorMessage string) model.Response {
	return model.Response{
		Success: isSuccess,
		Data:    data,
		Error:   errorMessage,
	}
}
