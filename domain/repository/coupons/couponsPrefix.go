package repository

import (
	entity "assigment-final-project/domain/entity/coupons"
	"context"
)

type CouponsPrefix interface {
	InsertPrefix(ctx context.Context, prefix *entity.CouponsPrefix) error
	InsertPrefixs(ctx context.Context, prefixs []*entity.CouponsPrefix) error
	GetPrefixs(ctx context.Context, offsetNum, limitNum int) ([]*entity.CouponsPrefix, error)
	GetPrefixMinimumPrice(ctx context.Context, price float64) ([]*entity.CouponsPrefix, error)
	UpdatePrefix(ctx context.Context, prefix *entity.CouponsPrefix) (*entity.CouponsPrefix, error)
	FindCouponPrefix(ctx context.Context, prefix, criteria string) (*entity.CouponsPrefix, error)
	DeletePrefix(ctx context.Context, id int) error
}
