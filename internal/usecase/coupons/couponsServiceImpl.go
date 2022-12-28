package usecase

import (
	repository "assigment-final-project/domain/repository/coupons"
	customers "assigment-final-project/domain/repository/customers"
	"assigment-final-project/helper"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
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

func (c *CouponServiceImpl) GetCouponByCustomerId(ctx context.Context, customerId string, page int) (*http_response.CouponsCustomerResponse, int, error) {
	var (
		limit        = 5
		offset       = limit * (page - 1)
		wg           = sync.WaitGroup{}
		chanCustomer = make(chan *http_response.CustomerResponse)
		coupons      = make([]*http_response.CouponsResponse, 0)
	)

	result, err := c.repoCoupon.FindCouponByCustomerId(ctx, customerId, offset, limit)
	helper.PanicIfError(err)
	if len(result) == 0 {
		return nil, 0, errors.New("coupon not found")
	}

	coupons = http_response.ListDomainCouponsToCouponsResponse(result)
	wg.Add(1)
	go func() {
		defer wg.Done()
		dataCustomer, err2 := c.repoCustomer.FindCustomerById(ctx, customerId, "")
		helper.PrintIfError(err2)
		chanCustomer <- http_response.DomainToCustomerResponse(dataCustomer)
	}()
	customerResult := <-chanCustomer
	wg.Wait()

	rows := helper.CountTotalRows(ctx, mysql_connection.InitMysqlDB(), "coupons")
	return &http_response.CouponsCustomerResponse{
		Customer: customerResult,
		Coupons:  coupons,
	}, rows.TotalRows, nil
}
