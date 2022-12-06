package repository

import (
	entity "assigment-final-project/domain/entity/products"
	"context"
)

type ProductRepo interface {
	InsertProducts(ctx context.Context, products *entity.Products) error
	FindProduct(ctx context.Context, productId string) (*entity.Products, error)
	GetProducts(ctx context.Context) ([]*entity.Products, error)
	UpdateProduct(ctx context.Context, product *entity.Products) (*entity.Products, error)
	DeleteProduct(ctx context.Context, productId string) error
}
