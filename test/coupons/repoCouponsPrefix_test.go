package coupons_test

import (
	entity "assigment-final-project/domain/entity/coupons"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	repository "assigment-final-project/internal/repository/mysql"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

var (
	ctx               = context.Background()
	db                = mysql_connection.InitMysqlDB()
	repoCouponsPrefix = repository.NewCouponPrefixImpl(db)
	repoCoupons       = repository.NewCouponsRepoImpl(db)
)

func TestInsertCouponsPrefix(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2022-10-11")

	err := repoCouponsPrefix.InsertPrefix(ctx, entity.NewCouponsPrefix(&entity.DTOCouponsPrefix{
		PrefixName:   "ulti",
		MinimumPrice: 100000,
		Discount:     30,
		ExpireDate:   date,
		Criteria:     "console",
		CreatedAt:    time.Now(),
	}))

	fmt.Println(err)
	assert.NoError(t, err)
}

func TestGetCouponsPrefix(t *testing.T) {
	prefixs, err := repoCouponsPrefix.GetPrefixs(ctx)
	fmt.Println(err, " this error line")
	for _, prefix := range prefixs {
		fmt.Println(prefix)
		fmt.Println(prefix.ExpireDate())
		fmt.Println(prefix.CreatedAt())
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, prefixs)
}

func TestUpdateCounponsPrefix(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2023-10-11")
	prefix, err := repoCouponsPrefix.UpdatePrefix(ctx, entity.NewCouponsPrefix(&entity.DTOCouponsPrefix{
		Id:           1,
		PrefixName:   "prime",
		MinimumPrice: 100000,
		Discount:     10,
		ExpireDate:   date,
		Criteria:     "new game",
	}))

	assert.NoError(t, err)
	assert.NotEmpty(t, prefix)
	assert.Equal(t, "prime", prefix.PrefixName())
}
func TestDeleteCouponsPrefix(t *testing.T) {
	err := repoCouponsPrefix.DeletePrefix(ctx, 1)
	assert.NoError(t, err)
}

func TestFindCouponPrefix(t *testing.T) {
	prefix, err := repoCouponsPrefix.FindCouponPrefix(ctx, "ulti", "computer")
	assert.NoError(t, err)
	assert.NotEmpty(t, prefix)
	assert.Equal(t, "ulti", prefix.PrefixName())
	fmt.Println(prefix)
}

func TestGetCoupons(t *testing.T) {
	result, err := repoCoupons.FindCouponByCustomerIdAndCode(ctx, "ULTI-pbFut594daUVgOZr", "bie7")
	assert.NoError(t, err)
	str := strings.Split(result.CouponCode(), "-")[0]
	fmt.Println(str)
}
