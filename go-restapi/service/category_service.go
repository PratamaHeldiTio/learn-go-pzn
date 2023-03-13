package service

import (
	"context"
	"go-restapi/model/request"
	"go-restapi/model/response"
)

// interface contrac service
type CategoryService interface {
	Created(ctx context.Context, request request.CategoryCreatedRequest) response.CategoryResponse
	Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
