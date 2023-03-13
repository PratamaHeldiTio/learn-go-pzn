package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-restapi/helper"
	requestModel "go-restapi/model/request"
	"go-restapi/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Created(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// make var for temp request
	var categoryCreatedRequest requestModel.CategoryCreatedRequest
	// process decoder
	helper.ReadFromRequestBody(request, &categoryCreatedRequest)

	// do controler with passing context and request
	categoryResponse := controller.CategoryService.Created(request.Context(), categoryCreatedRequest)

	// add reseponse from controler to api response
	APIResponse := helper.ResponseSuccess(categoryResponse)

	// process write response
	helper.WriteToResponseBody(writer, APIResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// decoder stream request json to golang type response
	var categoryUpdateRequest requestModel.CategoryUpdateRequest

	//process decoder
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	// get id from params and convert string to int
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)

	// add id to data request
	categoryUpdateRequest.Id = id

	// do controller with passing context and request
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)

	// add reseponse from controler to api response
	APIResponse := helper.ResponseSuccess(categoryResponse)

	// process write response
	helper.WriteToResponseBody(writer, APIResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// get id from params and convert string to int
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)

	// do controller with passing context and request
	controller.CategoryService.Delete(request.Context(), id)

	// add reseponse from controler to api response
	APIResponse := helper.ResponseSuccess(nil)

	// process write response
	helper.WriteToResponseBody(writer, APIResponse)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// get id from params and convert string to int
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicError(err)

	// do controller with passing context and request
	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	// add reseponse from controler to api response
	APIResponse := helper.ResponseSuccess(categoryResponse)

	// process write response
	helper.WriteToResponseBody(writer, APIResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// do controller with passing context and request
	categoryResponses := controller.CategoryService.FindAll(request.Context())

	// add reseponse from controler to api response
	APIResponse := helper.ResponseSuccess(categoryResponses)

	// process write response
	helper.WriteToResponseBody(writer, APIResponse)
}
