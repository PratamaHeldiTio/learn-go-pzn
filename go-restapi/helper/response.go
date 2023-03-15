package helper

import (
	"go-restapi/model/response"
)

func ResponseSuccess(data interface{}) response.APIResponseSucess {
	APIResponse := response.APIResponseSucess{
		Code:   200,
		Status: "Success",
		Data:   data,
	}
	return APIResponse
}

func ResponseError(code int, status string, err interface{}) (APIResponse response.APIResponseError) {
	APIResponse = response.APIResponseError{
		Code:    code,
		Status:  status,
		Message: err,
	}
	return
}
