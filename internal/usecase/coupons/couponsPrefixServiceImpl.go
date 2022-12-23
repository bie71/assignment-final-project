package usecase

import (
	repository "assigment-final-project/domain/repository/coupons"
	"assigment-final-project/helper"
	"assigment-final-project/helper/requestToEntity"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"os"
	"strconv"
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
	if !dateExpire.After(time.Now()) {
		return "", errors.New("expiration date must be later than the current date")
	}

	err = c.repoCouponsPrefix.InsertPrefix(ctx, requestToEntity.CouponPrefixRequestToDomainEntity(prefix, dateExpire))
	if err != nil {
		return "", err
	}
	return "Success Add Coupon", nil
}

func (c *CouponsPrefixServiceImpl) GetCoupons(ctx context.Context, page int) ([]*http_response.CouponsPrefixResponse, error) {
	var (
		limit, _ = strconv.Atoi(os.Getenv("LIMIT"))
		offset   = limit * (page - 1)
	)

	data, err := c.repoCouponsPrefix.GetPrefixs(ctx, offset, limit)
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
	if !dateExpire.After(time.Now()) {
		return nil, errors.New("expiration date must be later than the current date")
	}

	updatePrefix, err := c.repoCouponsPrefix.UpdatePrefix(ctx, requestToEntity.CouponPrefixRequestToDomainEntity(prefix, dateExpire))
	if err != nil {
		return nil, err
	}
	return http_response.DomainToCouponsPrefixResponse(updatePrefix), nil
}

func (c *CouponsPrefixServiceImpl) DeleteCoupon(ctx context.Context, id int) (string, error) {
	err := c.repoCouponsPrefix.DeletePrefix(ctx, id)
	if err != nil {
		return "", errors.New("coupon prefix not found")
	}
	return "Success Delete Coupon", nil
}
