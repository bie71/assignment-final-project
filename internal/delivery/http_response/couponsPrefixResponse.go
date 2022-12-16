package http_response

import (
	entity "assigment-final-project/domain/entity/coupons"
	"time"
)

type CouponsPrefixResponse struct {
	Id           int       `json:"id"`
	PrefixName   string    `json:"prefix_name" `
	MinimumPrice int64     `json:"minimum_price" `
	Discount     uint      `json:"discount" `
	ExpireDate   time.Time `json:"expire_date" `
	Criteria     string    `json:"criteria"`
	CreatedAt    time.Time `json:"created_at"`
}

func DomainToCouponsPrefixResponse(domain *entity.CouponsPrefix) *CouponsPrefixResponse {
	return &CouponsPrefixResponse{
		Id:           domain.Id(),
		PrefixName:   domain.PrefixName(),
		MinimumPrice: domain.MinimumPrice(),
		Discount:     domain.Discount(),
		ExpireDate:   domain.ExpireDate(),
		Criteria:     domain.Criteria(),
		CreatedAt:    domain.CreatedAt(),
	}
}

func ListDomainToListCouponsPrefixResponse(list []*entity.CouponsPrefix) []*CouponsPrefixResponse {
	listResponse := make([]*CouponsPrefixResponse, 0)

	for _, prefix := range list {
		result := DomainToCouponsPrefixResponse(prefix)
		listResponse = append(listResponse, result)
	}
	return listResponse
}
