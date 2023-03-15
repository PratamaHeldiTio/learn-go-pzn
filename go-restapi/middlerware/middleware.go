package middlerware

import (
	"go-restapi/helper"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		// make response format json
		APIResponse := helper.ResponseError(http.StatusUnauthorized, "Unauthorized Error", "Unauthorized")

		// write response
		helper.WriteToResponseBody(writer, APIResponse)
	}
}
