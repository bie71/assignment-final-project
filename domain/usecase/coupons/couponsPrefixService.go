package usecase

import (
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type CouponsPrefixService interface {
	AddCoupon(ctx context.Context, prefix *http_request.CouponsPrefixRequest) (string, error)
	GetCoupons(ctx context.Context) ([]*http_response.CouponsPrefixResponse, error)
	UpdateCoupon(ctx context.Context, prefix *http_request.CouponsPrefixRequest, id int) (*http_response.CouponsPrefixResponse, error)
	DeleteCoupon(ctx context.Context, id int) (string, error)
}
