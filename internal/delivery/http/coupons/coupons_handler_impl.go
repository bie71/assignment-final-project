package handler

import (
	usecase "assigment-final-project/domain/usecase/coupons"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type CouponHandlerImpl struct {
	serviceCoupons usecase.CouponsPrefixService
}

func NewCouponHandlerImpl(serviceCoupons usecase.CouponsPrefixService) *CouponHandlerImpl {
	return &CouponHandlerImpl{serviceCoupons: serviceCoupons}
}

func (c *CouponHandlerImpl) AddCoupon(w http.ResponseWriter, r *http.Request) {
	couponRequest := &http_request.CouponsPrefixRequest{}
	helper.ReadFromRequestBody(r, couponRequest)

	result, err := c.serviceCoupons.AddCoupon(r.Context(), couponRequest)
	if err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		delivery.ResponseDelivery(w, http.StatusBadRequest, nil, errors.Error())
		return
	}

	delivery.ResponseDelivery(w, http.StatusCreated, result, nil)
}

func (c *CouponHandlerImpl) GetCoupons(w http.ResponseWriter, r *http.Request) {
	result, err := c.serviceCoupons.GetCoupons(r.Context())
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, result, nil)
}

func (c *CouponHandlerImpl) UpdateAndDeleteCoupon(w http.ResponseWriter, r *http.Request) {
	query, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	if r.Method == http.MethodPut {
		couponRequest := &http_request.CouponsPrefixRequest{}
		helper.ReadFromRequestBody(r, couponRequest)

		result, err := c.serviceCoupons.UpdateCoupon(r.Context(), couponRequest, query)
		if err != nil {
			errors, ok := err.(validator.ValidationErrors)
			if !ok {
				delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
				return
			}

			delivery.ResponseDelivery(w, http.StatusBadRequest, nil, errors.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, result, nil)

	}

	result, err := c.serviceCoupons.DeleteCoupon(r.Context(), query)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, result, nil)
}
