package repository

import (
	entity "assigment-final-project/domain/entity/criteria"
	"context"
)

type CriteriaRepo interface {
	InsertCriteria(ctx context.Context, criteria *entity.Criteria) error
	GetCriteria(ctx context.Context) ([]*entity.Criteria, error)
	UpdateCriteria(ctx context.Context, criteria *entity.Criteria) (*entity.Criteria, error)
	DeleteCriteria(ctx context.Context, id int) error
}
