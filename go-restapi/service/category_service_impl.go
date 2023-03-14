package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-restapi/helper"
	"go-restapi/model/domain"
	"go-restapi/model/request"
	"go-restapi/model/response"
	"go-restapi/repository"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository // repository interface for connest service to repository
	DB         *sql.DB                       // for begin dan commit or rollback db and assign to argument repository
	Validate   *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Repository: categoryRepository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *CategoryServiceImpl) Created(ctx context.Context, request request.CategoryCreatedRequest) response.CategoryResponse {
	// do validation request before use with package validation
	// before validation must add tag validate in struct data want to validation
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	// open db transaction
	tx, err := service.DB.Begin()
	helper.PanicError(err)            // if err panic
	defer helper.CommitOrRollback(tx) // recovery func for rollback or commit db transaction

	// collect request to domain bcs domain for query db
	category := domain.Category{
		Name: request.Name,
	}

	// do action with db in repository
	category = service.Repository.Save(ctx, tx, category)

	// return with mapping data domain category to response category
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse {
	// do validation request before use with package validation
	// before validation must add tag validate in struct data want to validation
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	// membuka db transaction
	tx, err := service.DB.Begin()
	helper.PanicError(err)            // if err panic
	defer helper.CommitOrRollback(tx) // recovery func for rollback or commit db transaction

	// checking id is exist
	category, err := service.Repository.FindById(ctx, tx, request.Id)
	helper.PanicError(err)

	// set field update
	category.Name = request.Name

	// do action with db in repository
	category = service.Repository.Update(ctx, tx, category)

	// return with mapping data domain category to response category
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// membuka db transaction
	tx, err := service.DB.Begin()
	helper.PanicError(err)            // if err panic
	defer helper.CommitOrRollback(tx) // recovery func for rollback or commit db transaction

	// checking id is exist
	_, err = service.Repository.FindById(ctx, tx, categoryId)
	helper.PanicError(err)

	// do action with db in repository
	service.Repository.Delete(ctx, tx, categoryId)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) response.CategoryResponse {
	// membuka db transaction
	tx, err := service.DB.Begin()
	helper.PanicError(err)            // if err panic
	defer helper.CommitOrRollback(tx) // recovery func for rollback or commit db transaction

	// checking id is exist
	category, err := service.Repository.FindById(ctx, tx, categoryId)
	helper.PanicError(err)

	// return with mapping data domain category to response category
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	// membuka db transaction
	tx, err := service.DB.Begin()
	helper.PanicError(err)            // if err panic
	defer helper.CommitOrRollback(tx) // recovery func for rollback or commit db transaction

	// do action with db in repository
	categories := service.Repository.FindAll(ctx, tx)
	helper.PanicError(err)

	// return with mapping data domain category to response category
	return helper.ToCategoryResponses(categories)
}
