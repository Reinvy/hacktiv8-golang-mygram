package util

import "mygram/model/dto"

func CreateResponse(status string, message string, data any) dto.Response {
	return dto.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
