package usecase

import (
	entity "assigment-final-project/domain/entity/coupons"
	repository "assigment-final-project/domain/repository/coupons"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"github.com/go-playground/validator/v10"
	"time"
)

type CouponsPrefixServiceImpl struct {
	repoCouponsPrefix repository.CouponsPrefix
	validation        *validator.Validate
}

func NewCouponsPrefixServiceImpl(repoCouponsPrefix repository.CouponsPrefix, validation *validator.Validate) *CouponsPrefixServiceImpl {
	return &CouponsPrefixServiceImpl{repoCouponsPrefix: repoCouponsPrefix, validation: validation}
}

func (c *CouponsPrefixServiceImpl) AddCoupon(ctx context.Context, prefix *http_request.CouponsPrefixRequest) (string, error) {
	errValidation := c.validation.Struct(prefix)
	if errValidation != nil {
		return "", errValidation
	}

	dateExpire, err := time.Parse("2006-01-02", prefix.ExpireDate)
	helper.PanicIfError(err)

	data := entity.NewCouponsPrefix(&entity.DTOCouponsPrefix{
		PrefixName:   prefix.PrefixName,
		MinimumPrice: prefix.MinimumPrice,
		Discount:     prefix.Discount,
		ExpireDate:   dateExpire,
		Criteria:     prefix.Criteria,
		CreatedAt:    time.Now(),
	})

	err = c.repoCouponsPrefix.InsertPrefix(ctx, data)
	if err != nil {
		return "", err
	}
	return "Success Add Coupon", nil
}

func (c *CouponsPrefixServiceImpl) GetCoupons(ctx context.Context) ([]*http_response.CouponsPrefixResponse, error) {
	data, err := c.repoCouponsPrefix.GetPrefixs(ctx)
	if err != nil {
		return nil, err
	}
	return http_response.ListDomainToListCouponsPrefixResponse(data), nil
}

func (c *CouponsPrefixServiceImpl) UpdateCoupon(ctx context.Context, prefix *http_request.CouponsPrefixRequest, id int) (*http_response.CouponsPrefixResponse, error) {
	errValidation := c.validation.Struct(prefix)
	if errValidation != nil {
		return nil, errValidation
	}

	dateExpire, err := time.Parse("2006-01-02", prefix.ExpireDate)
	helper.PanicIfError(err)

	data := entity.NewCouponsPrefix(&entity.DTOCouponsPrefix{
		Id:           id,
		PrefixName:   prefix.PrefixName,
		MinimumPrice: prefix.MinimumPrice,
		Discount:     prefix.Discount,
		ExpireDate:   dateExpire,
		Criteria:     prefix.Criteria,
		CreatedAt:    time.Now(),
	})

	updatePrefix, err := c.repoCouponsPrefix.UpdatePrefix(ctx, data)
	if err != nil {
		return nil, err
	}
	return http_response.DomainToCouponsPrefixResponse(updatePrefix), nil
}

func (c *CouponsPrefixServiceImpl) DeleteCoupon(ctx context.Context, id int) (string, error) {

	err := c.repoCouponsPrefix.DeletePrefix(ctx, id)
	if err != nil {
		return "", err
	}
	return "Success Delete Coupon", nil
}
