package usecase

import (
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type ProductsService interface {
	AddProduct(ctx context.Context, productRequest *http_request.ProductsRequest) (string, error)
	FindProductById(ctx context.Context, productId string) (*http_response.ProductsResponse, error)
	GetProducts(ctx context.Context, page int) ([]*http_response.ProductsResponse, error)
	UpdateProduct(ctx context.Context, productRequest *http_request.ProductsRequest, productID string) (*http_response.ProductsResponse, error)
	DeleteProductById(ctx context.Context, productId string) (string, error)
}
