package entity

import (
	"errors"
	"time"
)

type Customers struct {
	customerId string
	name       string
	contact    string
	createdAt  time.Time
}

type DTOCustomers struct {
	CustomerId string
	Name       string
	Contact    string
	CreatedAt  time.Time
}

func NewCustomer(DTOCustomers *DTOCustomers) (*Customers, error) {
	if DTOCustomers.CustomerId == "" {
		return nil, errors.New("customerid required")
	}
	if DTOCustomers.Name == "" {
		return nil, errors.New("name required")
	}
	if DTOCustomers.Contact == "" {
		return nil, errors.New("contact required")
	}

	return &Customers{
		customerId: DTOCustomers.CustomerId,
		name:       DTOCustomers.Name,
		contact:    DTOCustomers.Contact,
		createdAt:  DTOCustomers.CreatedAt,
	}, nil
}

func CustomersFromDb(customers *DTOCustomers) *Customers {
	return &Customers{
		customerId: customers.CustomerId,
		name:       customers.Name,
		contact:    customers.Contact,
		createdAt:  customers.CreatedAt,
	}
}

func (c *Customers) CustomerId() string {
	return c.customerId
}

func (c *Customers) Name() string {
	return c.name
}

func (c *Customers) Contact() string {
	return c.contact
}

func (c *Customers) CreatedAt() time.Time {
	return c.createdAt
}
