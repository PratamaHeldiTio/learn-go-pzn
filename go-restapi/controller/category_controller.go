package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// CategoryController interface controller must have 3 parameter default in http handler and router
type CategoryController interface {
	Created(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
