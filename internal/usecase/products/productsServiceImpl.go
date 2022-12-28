package usecase

import (
	repository "assigment-final-project/domain/repository/products"
	"assigment-final-project/helper"
	"assigment-final-project/helper/requestToEntity"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"os"
	"strconv"
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

	dataProduct := requestToEntity.ProductRequestToEntity(productRequest, `product-`+helper.RandomString(16))
	err := p.repoProducts.InsertProduct(ctx, dataProduct)
	if err != nil {
		return "", err
	}

	return dataProduct.ProductId(), nil
}

func (p *ProductsServiceImpl) FindProductById(ctx context.Context, productId string) (*http_response.ProductsResponse, error) {
	product, err := p.repoProducts.FindProduct(ctx, productId)
	if err != nil || product == nil {
		return nil, errors.New("product not found")
	}

	return http_response.DomainProductsToProductsResponse(product), nil
}

func (p *ProductsServiceImpl) GetProducts(ctx context.Context, page int) ([]*http_response.ProductsResponse, int, error) {
	var (
		limit, _ = strconv.Atoi(os.Getenv("LIMIT"))
		offset   = limit * (page - 1)
	)
	products, err := p.repoProducts.GetProducts(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	rows := helper.CountTotalRows(ctx, mysql_connection.InitMysqlDB(), "products")
	return http_response.ListDomainProductToResponseProducts(products), rows.TotalRows, nil
}

func (p *ProductsServiceImpl) UpdateProduct(ctx context.Context, productRequest *http_request.ProductsRequest, productId string) (*http_response.ProductsResponse, error) {
	errValidation := p.validation.Struct(productRequest)
	if errValidation != nil {
		return nil, errValidation
	}
	result, err := p.repoProducts.FindProduct(ctx, productId)
	if err != nil || result == nil {
		return nil, errors.New("product not found")
	}

	dataProduct := requestToEntity.ProductRequestToEntity(productRequest, productId)
	product, err := p.repoProducts.UpdateProduct(ctx, dataProduct)
	helper.PanicIfError(err)

	return http_response.DomainProductsToProductsResponse(product), nil
}

func (p *ProductsServiceImpl) DeleteProductById(ctx context.Context, productId string) (string, error) {
	result, err := p.repoProducts.FindProduct(ctx, productId)
	if err != nil || result == nil {
		return "", errors.New("product not found")
	}
	err = p.repoProducts.DeleteProduct(ctx, productId)
	helper.PanicIfError(err)
	return "Success Delete Product", nil
}
