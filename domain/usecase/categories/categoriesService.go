package usecase

import (
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type CategoryService interface {
	AddCategory(ctx context.Context, categoryRequest *http_request.CategoryRequest) (string, error)
	FindCategoryById(ctx context.Context, categoryId string) (*http_response.CategoryResponse, error)
	GetCategories(ctx context.Context) ([]*http_response.CategoryResponse, error)
	DeleteCategoryById(ctx context.Context, categoryId string) (string, error)
}
