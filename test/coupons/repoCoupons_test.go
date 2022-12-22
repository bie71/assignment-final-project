package coupons

import (
	"assigment-final-project/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCouponsByCustomerId(t *testing.T) {
	result, err := repoCoupons.FindCouponByCustomerId(ctx, "123", 0, 10)
	helper.PrintIfError(err)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	for _, coupon := range result {
		fmt.Println(coupon)
	}

}
