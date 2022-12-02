package usecase

import (
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type CustomersService interface {
	AddCustomer(ctx context.Context, customerRequest *http_request.CustomerRequest) (string, error)
	FindCustomer(ctx context.Context, customerId, phone string) (*http_response.CustomerResponse, error)
	GetCustomers(ctx context.Context) ([]*http_response.CustomerResponse, error)
	DeleteCustomer(ctx context.Context, customerId, phone string) (string, error)
}
