package repository

import (
	entity "assigment-final-project/domain/entity/coupons"
	"context"
)

type CouponsPrefix interface {
	InsertPrefix(ctx context.Context, prefix *entity.CouponsPrefix) error
	GetPrefixs(ctx context.Context) ([]*entity.CouponsPrefix, error)
	UpdatePrefix(ctx context.Context, prefix *entity.CouponsPrefix) (*entity.CouponsPrefix, error)
	DeletePrefix(ctx context.Context, id int) error
}