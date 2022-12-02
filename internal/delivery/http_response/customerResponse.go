package http_response

import (
	entity "assigment-final-project/domain/entity/customers"
	"time"
)

type CustomerResponse struct {
	CustomerId string    `json:"customer_id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
}

func DomainToCustomerResponse(entity *entity.Customers) *CustomerResponse {
	return &CustomerResponse{
		CustomerId: entity.CustomerId(),
		Name:       entity.Name(),
		Phone:      entity.Contact(),
		CreatedAt:  entity.CreatedAt(),
	}
}

func ListDomainToListCustomerResponse(listEntity []*entity.Customers) []*CustomerResponse {
	listCustomer := make([]*CustomerResponse, 0)

	for _, customers := range listEntity {
		response := DomainToCustomerResponse(customers)
		listCustomer = append(listCustomer, response)
	}
	return listCustomer
}
