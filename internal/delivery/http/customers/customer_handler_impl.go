package customers

import (
	usecase "assigment-final-project/domain/usecase/customers"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"net/http"
)

type CustomerHandlerImpl struct {
	customerService usecase.CustomersService
}

func NewCustomerHandlerImpl(customerService usecase.CustomersService) *CustomerHandlerImpl {
	return &CustomerHandlerImpl{customerService: customerService}
}

func (c *CustomerHandlerImpl) AddCustomer(w http.ResponseWriter, r *http.Request) {
	customerRequest := &http_request.CustomerRequest{}
	helper.ReadFromRequestBody(r, customerRequest)

	data, err := c.customerService.AddCustomer(r.Context(), customerRequest)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusBadRequest, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusCreated, data)
	return
}

func (c *CustomerHandlerImpl) GetAndDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := r.URL.Query().Get("id")

	if r.Method == http.MethodDelete {
		data, err := c.customerService.DeleteCustomer(r.Context(), customerId, customerId)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusNotFound, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, data)
		return
	}

	if customerId == "" {
		data, err := c.customerService.GetCustomers(r.Context())
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, data)
		return
	}

	data, err := c.customerService.FindCustomer(r.Context(), customerId, customerId)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, data)
}
