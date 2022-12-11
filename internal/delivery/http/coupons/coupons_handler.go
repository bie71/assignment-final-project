package handler

import "net/http"

type CouponsHandler interface {
	AddCoupon(w http.ResponseWriter, r *http.Request)
	GetCoupons(w http.ResponseWriter, r *http.Request)
	UpdateAndDeleteCoupon(w http.ResponseWriter, r *http.Request)
}
