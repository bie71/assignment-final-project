package coupons_test

import (
	"assigment-final-project/internal/delivery/http_request"
	usecase "assigment-final-project/internal/usecase/coupons"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	validation           = validator.New()
	serviceCouponsPrefix = usecase.NewCouponsPrefixServiceImpl(repoCouponsPrefix, validation)
)

func TestAddCoupons(t *testing.T) {
	data := &http_request.CouponsPrefixRequest{
		PrefixName:   "ulti",
		MinimumPrice: 2000,
		Discount:     10,
		ExpireDate:   "2023-01-01",
		Criteria:     "new console",
	}
	result, err := serviceCouponsPrefix.AddCoupon(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestGetCouponsPrefixService(t *testing.T) {
	coupons, err := serviceCouponsPrefix.GetCoupons(ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, coupon := range coupons {
		fmt.Println(coupon)
		fmt.Println(coupon.Id)
		fmt.Println(coupon.PrefixName)
		fmt.Println(coupon.MinimumPrice)
		fmt.Println(coupon.ExpireDate)
		fmt.Println(coupon.Discount)
		fmt.Println(coupon.Criteria)
		fmt.Println(coupon.CreatedAt)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, coupons)
}

func TestUpdateCouponsPrefixService(t *testing.T) {
	data := &http_request.CouponsPrefixRequest{
		PrefixName:   "Prime",
		MinimumPrice: 3000,
		Discount:     20,
		ExpireDate:   "2023-11-01",
		Criteria:     "new game",
	}
	coupon, err := serviceCouponsPrefix.UpdateCoupon(ctx, data, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(coupon)

	assert.NoError(t, err)
	assert.NotEmpty(t, coupon)
}

func TestDeleteCouponsPrefixService(t *testing.T) {
	result, err := serviceCouponsPrefix.DeleteCoupon(ctx, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}
