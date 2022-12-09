package usecase

import (
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type CriteriaService interface {
	AddCriteria(ctx context.Context, criteriaRequest *http_request.CriteriaRequest) (string, error)
	GetCriteria(ctx context.Context) ([]*http_response.CriteriaResponse, error)
	UpdateCriteria(ctx context.Context, criteriaRequest *http_request.CriteriaRequest, id int) (*http_response.CriteriaResponse, error)
	DeleteCriteria(ctx context.Context, id int) (string, error)
}
