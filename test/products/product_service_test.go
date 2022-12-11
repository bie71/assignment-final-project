package products_test

import (
	"assigment-final-project/internal/delivery/http_request"
	usecase "assigment-final-project/internal/usecase/products"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var serviceProduct = usecase.NewProductsServiceImpl(repo, validation)

func TestAddProduct(t *testing.T) {
	dataRequest := &http_request.ProductsRequest{
		Name:       "ps 5",
		Price:      1500,
		CategoryId: "console",
		Stock:      20,
	}
	response, err := serviceProduct.AddProduct(ctx, dataRequest)
	if err != nil {
		fmt.Errorf("error in service %s", err)
	}
	fmt.Println(response)

	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}

func TestGetProductsService(t *testing.T) {
	products, err := serviceProduct.GetProducts(ctx)
	if err != nil {
		fmt.Errorf("error in service %s", err)
	}
	for _, product := range products {
		fmt.Println(product)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, products)
}

func TestFindProductById(t *testing.T) {
	product, err := serviceProduct.FindProductById(ctx, "product-ytRvt5O3Q7glxSww")
	if err != nil {
		fmt.Errorf("error in service %s", err)
	}
	fmt.Println(product)
	assert.NotEmpty(t, product)
	assert.NoError(t, err)
	assert.Equal(t, "product-ytRvt5O3Q7glxSww", product.ProductId)
	assert.Equal(t, "ps 5", product.Name)
}

func TestUpdateById(t *testing.T) {
	dataRequest := &http_request.ProductsRequest{
		Name:       "ps 5 updated",
		Price:      2000,
		CategoryId: "console updated",
		Stock:      30,
	}

	response, err := serviceProduct.UpdateProduct(ctx, dataRequest, "product-ytRvt5O3Q7glxSww")
	if err != nil {
		fmt.Errorf("error in service %s", err)
	}
	fmt.Println(response)

	assert.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, "ps 5 updated", response.Name)
	assert.Equal(t, "console updated", response.CategoryId)
}

func TestDeleteById(t *testing.T) {
	response, err := serviceProduct.DeleteProductById(ctx, "product-ytRvt5O3Q7glxSww")
	if err != nil {
		fmt.Errorf("error in service %s", err)
	}
	fmt.Println(response)
	assert.NoError(t, err)
}
