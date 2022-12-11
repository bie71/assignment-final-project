package repository

import (
	entity "assigment-final-project/domain/entity/coupons"
	"context"
)

type CouponsRepo interface {
	InsertCoupon(ctx context.Context, coupons *entity.Coupons) error
	FindCoupon(ctx context.Context, code string, id int) (*entity.Coupons, error)
	GetCoupons(ctx context.Context) ([]*entity.Coupons, error)
	DeleteCoupon(ctx context.Context, code string, id int) error
}
