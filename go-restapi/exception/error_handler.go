package exception

import (
	"github.com/go-playground/validator/v10"
	"go-restapi/helper"
	"net/http"
)

// ErrorHandler error handler for response if have error
func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		// make response format json
		APIResponse := helper.ResponseError(http.StatusNotFound, "Not Found Error", exception.Error)

		// write response
		helper.WriteToResponseBody(writer, APIResponse)

		return true
	} else {
		return false
	}

}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		// make response format json
		APIResponse := helper.ResponseError(http.StatusBadRequest, "Bad Request Error", exception.Error())

		// write response
		helper.WriteToResponseBody(writer, APIResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	// set header
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	// make response format json
	APIResponse := helper.ResponseError(http.StatusInternalServerError, "Internal Server Error", err)

	// write response
	helper.WriteToResponseBody(writer, APIResponse)
}
