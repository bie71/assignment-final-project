package repository

import (
	entity "assigment-final-project/domain/entity/categories"
	"context"
)

type CategoryRepo interface {
	InsertCategory(ctx context.Context, category *entity.Categories) error
	InsertListCategory(ctx context.Context, categories []*entity.Categories) error
	FindCategory(ctx context.Context, categoryId string) (*entity.Categories, error)
	GetCategories(ctx context.Context, offsetNum, limitNum int) ([]*entity.Categories, error)
	DeleteCategory(ctx context.Context, categoryId string) error
}
