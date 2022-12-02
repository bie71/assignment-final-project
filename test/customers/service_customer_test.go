package customers_test

import (
	"assigment-final-project/internal/delivery/http_request"
	usecase "assigment-final-project/internal/usecase/customers"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	validation      = validator.New()
	serviceCustomer = usecase.NewCustomerServiceImpl(repoCustomer, validation)
)

func TestAddCustomer(t *testing.T) {
	customer, err := serviceCustomer.AddCustomer(ctx, &http_request.CustomerRequest{
		Name:  "habibi",
		Phone: "12234",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, customer)
}

func TestFindCustomer(t *testing.T) {
	customer, err := serviceCustomer.FindCustomer(ctx, "", "1234")
	assert.NoError(t, err)
	assert.NotEmpty(t, customer)
}

func TestDeleteCustomer(t *testing.T) {
	customer, err := serviceCustomer.DeleteCustomer(ctx, "", "1234")
	assert.NoError(t, err)
	assert.Equal(t, "Success Delete Customer", customer)
}

func TestGetCustomersService(t *testing.T) {
	customers, err := serviceCustomer.GetCustomers(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, customers)
}
