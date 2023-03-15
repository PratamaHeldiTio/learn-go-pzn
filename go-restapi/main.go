package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"go-restapi/app"
	"go-restapi/controller"
	"go-restapi/exception"
	"go-restapi/helper"
	"go-restapi/middlerware"
	"go-restapi/repository"
	"go-restapi/service"
	"net/http"
)

func main() {
	// open conection db
	db := app.NewDB()

	// define new validator
	validate := validator.New()

	// call repository, category service, controller
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// define new router
	router := httprouter.New()

	// htpp methode handler
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Created)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// panic handler, if ada panic program tidak berhenti melainkan masuk kseini
	router.PanicHandler = exception.ErrorHandler
	// define server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middlerware.NewAuthMiddleware(router),
	}

	//run server
	err := server.ListenAndServe()
	helper.PanicError(err)
}
