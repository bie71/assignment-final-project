package coupons

import (
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	repository "assigment-final-project/internal/repository/mysql"
	usecase "assigment-final-project/internal/usecase/coupons"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	db            = mysql_connection.InitMysqlDB()
	ctx           = context.Background()
	repoCoupons   = repository.NewCouponsRepoImpl(db)
	repoCustomer  = repository.NewCustomerRepoImpl(db)
	serviceCoupon = usecase.NewCouponServiceImpl(repoCoupons, repoCustomer)
)

func TestGetCouponByCustomerId(t *testing.T) {
	result, err := serviceCoupon.GetCouponByCustomerId(ctx, "bie7", 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	fmt.Println(result.Customer)
	for _, coupon := range result.Coupons {
		fmt.Println(coupon)
	}
}
