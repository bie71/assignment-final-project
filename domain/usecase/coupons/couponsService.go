package usecase

import (
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type CouponsService interface {
	GetCouponByCustomerId(ctx context.Context, customerId string, page int) (*http_response.CouponsCustomerResponse, error)
}
