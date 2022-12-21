package usecase

import (
	repository "assigment-final-project/domain/repository/coupons"
	customers "assigment-final-project/domain/repository/customers"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"sync"
)

type CouponServiceImpl struct {
	repoCoupon   repository.CouponsRepo
	repoCustomer customers.RepoCustomer
}

func NewCouponServiceImpl(repoCoupon repository.CouponsRepo, repoCustomer customers.RepoCustomer) *CouponServiceImpl {
	return &CouponServiceImpl{repoCoupon: repoCoupon, repoCustomer: repoCustomer}
}

func (c *CouponServiceImpl) GetCouponByCustomerId(ctx context.Context, customerId string, page int) (*http_response.CouponsCustomerResponse, error) {
	var (
		limit        = 10
		offset       = limit * (page - 1)
		wg           = sync.WaitGroup{}
		chanCustomer = make(chan *http_response.CustomerResponse)
		coupons      = make([]*http_response.CouponsResponse, 0)
	)

	result, err := c.repoCoupon.FindCouponByCustomerId(ctx, customerId, offset, limit)
	if result == nil && err != nil {
		return nil, errors.New("coupon not found")
	}
	for _, coupon := range result {
		dataCoupons := &http_response.CouponsResponse{
			CouponCode: coupon.CouponCode(),
			IsUsed:     coupon.IsUsed(),
			ExpireDate: coupon.ExpireDate(),
		}
		coupons = append(coupons, dataCoupons)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		dataCustomer, err2 := c.repoCustomer.FindCustomerById(ctx, customerId, "")
		helper.PrintIfError(err2)
		chanCustomer <- &http_response.CustomerResponse{
			CustomerId: dataCustomer.CustomerId(),
			Name:       dataCustomer.Name(),
			Phone:      dataCustomer.Contact(),
			CreatedAt:  dataCustomer.CreatedAt(),
		}
	}()
	customerResult := <-chanCustomer
	wg.Wait()

	return &http_response.CouponsCustomerResponse{
		Customer: customerResult,
		Coupons:  coupons,
	}, nil
}
