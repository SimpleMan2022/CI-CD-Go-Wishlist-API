package helper

import "go-wishlist-api-2/dto"

type ResponseWithData struct {
	Status  bool
	Code    int
	Message string
	Data    any
}

type ResponseWithError struct {
	Status  bool
	Code    int
	Message string
}

func Response(param dto.ResponseParam) any {
	var status bool
	var response any

	if param.StatusCode >= 200 && param.StatusCode < 300 {
		status = true
	} else {
		status = false
	}

	if param.Data != nil {
		response = ResponseWithData{
			Status:  status,
			Code:    param.StatusCode,
			Message: param.Message,
			Data:    param.Data,
		}
	} else {
		response = ResponseWithError{
			Status:  false,
			Code:    param.StatusCode,
			Message: param.Message,
		}
	}
	return response
}
