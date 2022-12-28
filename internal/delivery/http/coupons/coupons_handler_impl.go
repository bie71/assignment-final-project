package handler

import (
	usecase "assigment-final-project/domain/usecase/coupons"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

type CouponHandlerImpl struct {
	serviceCoupons usecase.CouponsPrefixService
	coupons        usecase.CouponsService
}

func NewCouponHandlerImpl(serviceCoupons usecase.CouponsPrefixService, coupons usecase.CouponsService) *CouponHandlerImpl {
	return &CouponHandlerImpl{serviceCoupons: serviceCoupons, coupons: coupons}
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
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, err := strconv.Atoi(page)
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	helper.PanicIfError(err)
	result, rows, err := c.serviceCoupons.GetCoupons(r.Context(), p)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, http_response.PaginationInfo(p, limit, rows, result), nil)
}

func (c *CouponHandlerImpl) UpdateAndDeleteCoupon(w http.ResponseWriter, r *http.Request) {
	query, err := strconv.Atoi(r.URL.Query().Get("id"))
	helper.PanicIfError(err)

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

func (c *CouponHandlerImpl) GetCouponsCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, err := strconv.Atoi(page)
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	helper.PanicIfError(err)
	result, rows, err := c.coupons.GetCouponByCustomerId(r.Context(), params["customerid"], p)
	if result == nil && err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, http_response.PaginationInfo(p, limit, rows, result), nil)
}
