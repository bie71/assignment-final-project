package customers_test

import (
	entity "assigment-final-project/domain/entity/customers"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	repository "assigment-final-project/internal/repository/mysql"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	db           = mysql_connection.InitMysqlDB()
	ctx          = context.Background()
	repoCustomer = repository.NewCustomerRepoImpl(db)
)

func TestInsertCustomer(t *testing.T) {
	data, _ := entity.NewCustomer(&entity.DTOCustomers{
		CustomerId: "123",
		Name:       "Habibi",
		Contact:    "08777123123",
		CreatedAt:  time.Now(),
	})
	err := repoCustomer.InsertCustomer(ctx, data)
	assert.NoError(t, err)
}

func TestFindByIdOrPhone(t *testing.T) {
	customer, err := repoCustomer.FindCustomerById(ctx, "123", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, customer)
}

func TestDeleteCustomerByIdOrPhone(t *testing.T) {
	err := repoCustomer.DeleteCustomerById(ctx, "123", "")
	assert.NoError(t, err)
}

func TestGetCustomers(t *testing.T) {
	customers, err := repoCustomer.GetCustomers(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, customers)
}
