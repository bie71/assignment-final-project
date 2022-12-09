package usecase

import (
	entity "assigment-final-project/domain/entity/products"
	repository "assigment-final-project/domain/repository/products"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"github.com/go-playground/validator/v10"
)

type ProductsServiceImpl struct {
	repoProducts repository.ProductRepo
	validation   *validator.Validate
}

func NewProductsServiceImpl(repoProducts repository.ProductRepo, validation *validator.Validate) *ProductsServiceImpl {
	return &ProductsServiceImpl{repoProducts: repoProducts, validation: validation}
}

func (p *ProductsServiceImpl) AddProduct(ctx context.Context, productRequest *http_request.ProductsRequest) (string, error) {
	errValidation := p.validation.Struct(productRequest)
	if errValidation != nil {
		return "", errValidation
	}
	productId := `product-` + helper.RandomString(16)

	dataProduct := entity.NewProducts(&entity.DTOProducts{
		ProductId:  productId,
		Name:       productRequest.Name,
		Price:      productRequest.Price,
		CategoryId: productRequest.CategoryId,
		Stock:      productRequest.Stock,
	})

	err := p.repoProducts.InsertProducts(ctx, dataProduct)
	if err != nil {
		return "", err
	}

	return dataProduct.ProductId(), nil
}

func (p *ProductsServiceImpl) FindProductById(ctx context.Context, productId string) (*http_response.ProductsResponse, error) {
	product, err := p.repoProducts.FindProduct(ctx, productId)
	if err != nil {
		return nil, err
	}

	return http_response.DomainProductsToProductsResponse(product), nil
}

func (p *ProductsServiceImpl) GetProducts(ctx context.Context) ([]*http_response.ProductsResponse, error) {
	products, err := p.repoProducts.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	return http_response.ListDomainProductToResponseProducts(products), nil
}

func (p *ProductsServiceImpl) UpdateProduct(ctx context.Context, productRequest *http_request.ProductsRequest, productId string) (*http_response.ProductsResponse, error) {
	errValidation := p.validation.Struct(productRequest)
	if errValidation != nil {
		return nil, errValidation
	}
	_, err := p.repoProducts.FindProduct(ctx, productId)
	if err != nil {
		return nil, err
	}

	dataProduct := entity.NewProducts(&entity.DTOProducts{
		ProductId:  productId,
		Name:       productRequest.Name,
		Price:      productRequest.Price,
		CategoryId: productRequest.CategoryId,
		Stock:      productRequest.Stock,
	})

	product, err := p.repoProducts.UpdateProduct(ctx, dataProduct)
	if err != nil {
		return nil, err
	}

	return http_response.DomainProductsToProductsResponse(product), nil
}

func (p *ProductsServiceImpl) DeleteProductById(ctx context.Context, productId string) (string, error) {
	_, err := p.repoProducts.FindProduct(ctx, productId)
	if err != nil {
		return "", err
	}

	err = p.repoProducts.DeleteProduct(ctx, productId)
	if err != nil {
		return "", err
	}
	return "Success Delete Product", nil
}
