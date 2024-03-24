package util

import "mygram/domain/dto"

func CreateResponse(status string, message string, data any) dto.Response {
	return dto.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
