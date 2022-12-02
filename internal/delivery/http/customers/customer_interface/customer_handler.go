package customer_interface

import "net/http"

type CustomerHandler interface {
	AddCustomer(w http.ResponseWriter, r *http.Request)
	GetAndDeleteCustomer(w http.ResponseWriter, r *http.Request)
}
