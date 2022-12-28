package customers

import (
	usecase "assigment-final-project/domain/usecase/customers"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"net/http"
	"os"
	"strconv"
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
		delivery.ResponseDelivery(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusCreated, data, nil)
}

func (c *CustomerHandlerImpl) GetAndDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerId := r.URL.Query().Get("id")
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, err := strconv.Atoi(page)
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	helper.PanicIfError(err)

	if r.Method == http.MethodDelete {
		data, err := c.customerService.DeleteCustomer(r.Context(), customerId, customerId)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, data, nil)
		return
	}

	if customerId == "" {
		data, rows, err := c.customerService.GetCustomers(r.Context(), p)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		delivery.ResponseDelivery(w, http.StatusOK, http_response.PaginationInfo(p, limit, rows, data), nil)
		return
	}

	data, err := c.customerService.FindCustomer(r.Context(), customerId, customerId)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, data, nil)
}
