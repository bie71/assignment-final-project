package models

import "time"

type CustomerModels struct {
	CustomerId string    `dbq:"customer_id"`
	Name       string    `dbq:"name"`
	Contact    string    `dbq:"contact"`
	CreatedAt  time.Time `dbq:"created_at"`
}

func TableNameCustomer() string {
	return "customers"
}

func FieldNameCustomers() []string {
	return []string{
		"customer_id",
		"name",
		"contact",
		"created_at",
	}
}
