package usecase

import (
	entity "assigment-final-project/domain/entity/categories"
	repository "assigment-final-project/domain/repository/categories"
	"assigment-final-project/helper"
	"assigment-final-project/helper/requestToEntity"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"os"
	"strconv"
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

	dataCategory := requestToEntity.CategoryRequestToEntity(categoryRequest, `category-`+helper.RandomString(16))
	err := c.repoCategory.InsertCategory(ctx, dataCategory)
	if err != nil {
		return "", err
	}
	return dataCategory.CategoryId(), nil
}

func (c *CategoryServiceImpl) FindCategoryById(ctx context.Context, categoryId string) (*http_response.CategoryResponse, error) {
	category, err := c.repoCategory.FindCategory(ctx, categoryId)
	if err != nil || category == nil {
		return nil, errors.New("category not found")
	}
	return http_response.DomainToCategoryResponse(category), nil
}

func (c *CategoryServiceImpl) GetCategories(ctx context.Context, page int) ([]*http_response.CategoryResponse, error) {
	var (
		limit, _ = strconv.Atoi(os.Getenv("LIMIT"))
		offset   = limit * (page - 1)
	)
	categories, err := c.repoCategory.GetCategories(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return http_response.ListDomainToListCategory(categories), nil
}

func (c *CategoryServiceImpl) DeleteCategoryById(ctx context.Context, categoryId string) (string, error) {
	category, err := c.repoCategory.FindCategory(ctx, categoryId)
	if err != nil || category == nil {
		return "", errors.New("category not found")
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
			listEntity = append(listEntity, requestToEntity.CategoryRequestToEntity(category, `category-`+helper.RandomString(16)))
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
