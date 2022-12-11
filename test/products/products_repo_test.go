package products_test

import (
	entity "assigment-final-project/domain/entity/products"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	repository "assigment-final-project/internal/repository/mysql"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	ctx        = context.Background()
	db         = mysql_connection.InitMysqlDB()
	validation = validator.New()
	repo       = repository.NewProductsRepoImpl(db)
)

func TestInsertProduct(t *testing.T) {
	dataProduct := entity.NewProducts(&entity.DTOProducts{
		ProductId:  "1234",
		Name:       "handphone",
		Price:      500,
		CategoryId: "123",
		Stock:      10,
	})

	err := repo.InsertProducts(ctx, dataProduct)
	if err != nil {
		fmt.Errorf("err in repo %s", err)
	}

	assert.NoError(t, err)
}

func TestFindProductbyId(t *testing.T) {
	product, err := repo.FindProduct(ctx, "123")
	if err != nil {
		fmt.Errorf("err in repo %s", err)
	}
	fmt.Println(product)

	assert.NoError(t, err)
	assert.NotEmpty(t, product)
}

func TestGetProducts(t *testing.T) {
	products, err := repo.GetProducts(ctx)
	if err != nil {
		fmt.Errorf("err in repo %s", err)
	}

	for _, product := range products {
		fmt.Println(product)
	}
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
}

func TestDeleteProduct(t *testing.T) {
	err := repo.DeleteProduct(ctx, "1234")
	if err != nil {
		fmt.Errorf("error in repo %s", err)
	}
	assert.NoError(t, err)
}

func TestUpdateProduct(t *testing.T) {
	dataProduct := entity.NewProducts(&entity.DTOProducts{
		ProductId:  "123",
		Name:       "handphone updated",
		Price:      200,
		CategoryId: "gadget updated",
		Stock:      10,
	})

	productUpdate, err := repo.UpdateProduct(ctx, dataProduct)
	if err != nil {
		fmt.Errorf("error in repo %s", err)
	}
	fmt.Println(productUpdate)

	assert.NotEmpty(t, productUpdate)
	assert.NoError(t, err)
}
