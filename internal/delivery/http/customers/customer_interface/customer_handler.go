package customer_interface

import "net/http"

type CustomerHandler interface {
	AddCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}
