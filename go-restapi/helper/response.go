package helper

import "go-restapi/model/response"

func ResponseSuccess(data interface{}) (APIResponse response.APIResponse) {
	if data != nil {
		APIResponse = response.APIResponse{
			Code:   200,
			Status: "Success",
			Data:   data,
		}
	} else {
		APIResponse = response.APIResponse{
			Code:   200,
			Status: "Success",
		}
	}

	return
}
