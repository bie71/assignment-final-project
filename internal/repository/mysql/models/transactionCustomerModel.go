package models

import "time"

type TransactionCustomer struct {
	TransactionId string    `dbq:"transaction_id"`
	CustomerId    string    `dbq:"customer_id"`
	Name          string    `dbq:"name"`
	Contact       string    `dbq:"contact"`
	CreatedAt     time.Time `dbq:"created_at"`
}
