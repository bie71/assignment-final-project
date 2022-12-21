package http_response

import (
	coupon "assigment-final-project/domain/entity/coupons"
	"time"
)

type CouponsCustomerResponse struct {
	Customer *CustomerResponse
	Coupons  []*CouponsResponse
}

type CouponsResponse struct {
	CouponCode string    `json:"coupon_code,omitempty"`
	IsUsed     bool      `json:"is_used,omitempty"`
	ExpireDate time.Time `json:"expire_date"`
}

func ListDomainCouponsToCouponsResponse(domain []*coupon.Coupons) []*CouponsResponse {
	coupons := make([]*CouponsResponse, 0)
	for _, d := range domain {
		coupon := &CouponsResponse{
			CouponCode: d.CouponCode(),
			IsUsed:     d.IsUsed(),
			ExpireDate: d.ExpireDate(),
		}
		coupons = append(coupons, coupon)
	}
	return coupons
}
