package repository

import (
	entity "assigment-final-project/domain/entity/customers"
	"context"
)

type RepoCustomer interface {
	InsertCustomer(ctx context.Context, customer *entity.Customers) error
	FindCustomerById(ctx context.Context, customerId, phone string) (*entity.Customers, error)
	GetCustomers(ctx context.Context) ([]*entity.Customers, error)
	DeleteCustomerById(ctx context.Context, customerId, phone string) error
}
