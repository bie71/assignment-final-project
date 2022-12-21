package repository

import (
	entity "assigment-final-project/domain/entity/coupons"
	"context"
)

type CouponsRepo interface {
	InsertCoupon(ctx context.Context, coupons *entity.Coupons) error
	FindCoupon(ctx context.Context, code string) (*entity.Coupons, error)
	FindCouponByCustomerIdAndCode(ctx context.Context, code, customerId string) (*entity.Coupons, error)
	FindCouponByCustomerId(ctx context.Context, customerId string, offsetNum, limitNum int) ([]*entity.Coupons, error)
	GetCoupons(ctx context.Context, offsetNum, limitNum int) ([]*entity.Coupons, error)
	DeleteCoupon(ctx context.Context, code string) error
	UpdateStatusCoupon(ctx context.Context, code, customerId string) (bool, error)
}
