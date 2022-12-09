package usecase

import (
	entity "assigment-final-project/domain/entity/categories"
	repository "assigment-final-project/domain/repository/categories"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"github.com/go-playground/validator/v10"
	"sync"
)

type CategoryServiceImpl struct {
	repoCategory repository.CategoryRepo
	validation   *validator.Validate
}

func NewCategoryServiceImpl(repoCategory repository.CategoryRepo, validation *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{repoCategory: repoCategory, validation: validation}
}

func (c *CategoryServiceImpl) AddCategory(ctx context.Context, categoryRequest *http_request.CategoryRequest) (string, error) {
	errValidation := c.validation.Struct(categoryRequest)
	if errValidation != nil {
		return "", errValidation
	}

	categoryId := `category-` + helper.RandomString(16)

	dataCategory := entity.NewCategories(&entity.DTOCategories{
		CategoryId: categoryId,
		Name:       categoryRequest.Name,
	})

	err := c.repoCategory.InsertCategory(ctx, dataCategory)
	if err != nil {
		return "", err
	}
	return dataCategory.CategoryId(), nil
}

func (c *CategoryServiceImpl) FindCategoryById(ctx context.Context, categoryId string) (*http_response.CategoryResponse, error) {
	category, err := c.repoCategory.FindCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return http_response.DomainToCategoryResponse(category), nil
}

func (c *CategoryServiceImpl) GetCategories(ctx context.Context) ([]*http_response.CategoryResponse, error) {
	categories, err := c.repoCategory.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return http_response.ListDomainToListCategory(categories), nil
}

func (c *CategoryServiceImpl) DeleteCategoryById(ctx context.Context, categoryId string) (string, error) {
	category, err := c.repoCategory.FindCategory(ctx, categoryId)
	if err != nil || category == nil {
		return "", err
	}
	err = c.repoCategory.DeleteCategory(ctx, categoryId)
	if err != nil {
		return "", err
	}
	return "Success Delete Category", nil
}

func (c *CategoryServiceImpl) AddCategories(ctx context.Context, categories []*http_request.CategoryRequest) (string, error) {
	wg := sync.WaitGroup{}
	err := make(chan error)

	for _, categoryReq := range categories {
		errValidation := c.validation.Struct(categoryReq)
		if errValidation != nil {
			return "", errValidation
		}
	}

	wg.Add(1)
	defer close(err)
	go func() {
		defer wg.Done()
		listEntity := make([]*entity.Categories, 0)

		for _, category := range categories {

			categoryId := `category-` + helper.RandomString(16)
			entity := entity.NewCategories(&entity.DTOCategories{
				CategoryId: categoryId,
				Name:       category.Name,
			})
			listEntity = append(listEntity, entity)
		}
		errRepo := c.repoCategory.InsertListCategory(ctx, listEntity)
		err <- errRepo
	}()
	resultErr := <-err
	wg.Wait()

	if resultErr != nil {
		return "", resultErr
	}
	return "Success Insert Categories", nil
}
