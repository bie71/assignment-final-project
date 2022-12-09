package usecase

import (
	entity "assigment-final-project/domain/entity/criteria"
	repository "assigment-final-project/domain/repository/criteria"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"github.com/go-playground/validator/v10"
)

type CriteriaServiceImpl struct {
	repoCriteria repository.CriteriaRepo
	validation   *validator.Validate
}

func (c *CriteriaServiceImpl) AddCriteria(ctx context.Context, criteriaRequest *http_request.CriteriaRequest) (string, error) {
	errValidation := c.validation.Struct(criteriaRequest)
	if errValidation != nil {
		return "", errValidation
	}

	data := entity.NewCriteria(&entity.DTOCriteria{
		Name: criteriaRequest.Name,
	})

	err := c.repoCriteria.InsertCriteria(ctx, data)
	if err != nil {
		return "", err
	}
	return "Success Insert Criteria", nil
}

func (c *CriteriaServiceImpl) GetCriteria(ctx context.Context) ([]*http_response.CriteriaResponse, error) {
	data, err := c.repoCriteria.GetCriteria(ctx)
	if err != nil {
		return nil, err
	}
	result := http_response.ListDomainToListCriteriaResponse(data)
	return result, nil
}

func (c *CriteriaServiceImpl) UpdateCriteria(ctx context.Context, criteriaRequest *http_request.CriteriaRequest, id int) (*http_response.CriteriaResponse, error) {
	errValidation := c.validation.Struct(criteriaRequest)
	if errValidation != nil {
		return nil, errValidation
	}

	data := entity.NewCriteria(&entity.DTOCriteria{
		Id:   id,
		Name: criteriaRequest.Name,
	})

	result, err := c.repoCriteria.UpdateCriteria(ctx, data)
	if err != nil {
		return nil, err
	}

	return http_response.DomainCriteriaToCriteriaResponse(result), nil
}

func (c *CriteriaServiceImpl) DeleteCriteria(ctx context.Context, id int) (string, error) {
	err := c.repoCriteria.DeleteCriteria(ctx, id)
	if err != nil {
		return "", err
	}
	return "Success Delete Criteria", nil
}
