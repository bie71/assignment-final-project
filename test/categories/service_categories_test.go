package categories_test

import (
	"assigment-final-project/internal/delivery/http_request"
	usecase "assigment-final-project/internal/usecase/categories"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var serviceCategories = usecase.NewCategoryServiceImpl(repoCategories, validator.New())

func TestInsertCategoryService(t *testing.T) {
	req := &http_request.CategoryRequest{Name: "console"}

	_, err := serviceCategories.AddCategory(ctx, req)
	if err != nil {
		fmt.Errorf("error in service %s", err)
	}

	assert.NoError(t, err)
}

func TestInsertCategories(t *testing.T) {
	sliceReq := make([]*http_request.CategoryRequest, 0)
	for i := 0; i < 10; i++ {
		req := &http_request.CategoryRequest{Name: "console"}
		sliceReq = append(sliceReq, req)
	}
	results, err := serviceCategories.AddCategories(ctx, sliceReq)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
}

func TestGetCategoriesService(t *testing.T) {
	categories, err := serviceCategories.GetCategories(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, category := range categories {
		fmt.Println(category)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, categories)
}

func TestFindCategory(t *testing.T) {
	byId, err := serviceCategories.FindCategoryById(ctx, "category-0wEgJdS0BCfdE9M5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(byId)
	assert.NoError(t, err)
	assert.NotEmpty(t, byId)
	assert.Equal(t, "console", byId.Name)
}

func TestDeleteCategoryService(t *testing.T) {
	result, err := serviceCategories.DeleteCategoryById(ctx, "category-0wEgJdS0BCfdE9M5")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	assert.NoError(t, err)
}
