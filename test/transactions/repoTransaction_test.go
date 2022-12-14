package transactions_test

import (
	entity "assigment-final-project/domain/entity/transactions"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	repository2 "assigment-final-project/internal/repository/mysql"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	db              = mysql_connection.InitMysqlDB()
	ctx             = context.Background()
	repoTransaction = repository2.NewTransactionRepoImpl(db)
)

func TestGetProductAndCategory(t *testing.T) {
	result, err := repoTransaction.GetProductJoinCategory(ctx, "p1")

	assert.NoError(t, err)
	fmt.Println(result)
	fmt.Println(result.CategoryId)
}

func TestCreateTransaction(t *testing.T) {
	err := repoTransaction.CreateTransaction(ctx, entity.NewTransaction(&entity.DTOTransaction{
		TransactionId: "1",
		CustomerId:    "bie7",
		CouponCode:    "",
		PurchaseDate:  time.Now(),
	}))
	assert.NoError(t, err)
}

func TestGetItemsProduct(t *testing.T) {
	product, err := repoTransaction.GetItemsProduct(ctx, "transaction-37bsGLqNIenuLOxZ")
	assert.NoError(t, err)
	assert.NotEmpty(t, product)
	for _, itemsProduct := range product {
		assert.NotEmpty(t, itemsProduct)
	}
}

func TestGetTransactionCustomer(t *testing.T) {
	customers, err := repoTransaction.GetTransactionCustomers(ctx, "transaction-xQPDmpLPLBIQqXs1")
	assert.NoError(t, err)
	assert.NotEmpty(t, customers)
	assert.Equal(t, "bie7", customers.CustomerId)
	fmt.Println(customers)
}

func TestDeleteTransactions(t *testing.T) {
	err := repoTransaction.DeleteTransaction(ctx, "transaction-xQPDmpLPLBIQqXs1")
	assert.NoError(t, err)
}

func TestFindTransactions(t *testing.T) {
	result, err := repoTransaction.FindTransaction(ctx, "transaction-vE6iHV0cpUJKL30H")
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}
