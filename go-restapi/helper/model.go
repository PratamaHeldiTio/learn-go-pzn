package helper

import (
	"go-restapi/model/domain"
	"go-restapi/model/response"
)

func ToCategoryResponse(category domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []response.CategoryResponse {
	var categoriesResponse []response.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, ToCategoryResponse(category))
	}

	return categoriesResponse
}
