package transactions_test

import (
	"assigment-final-project/internal/delivery/http_request"
	repoItems "assigment-final-project/internal/repository/mysql"
	usecase "assigment-final-project/internal/usecase/transactions"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	repoItem           = repoItems.NewTransactionItemsRepoImpl(db)
	initialCoupons     = repoItems.NewCouponPrefixImpl(db)
	coupons            = repoItems.NewCouponsRepoImpl(db)
	transactionService = usecase.NewTransactionServiceImpl(repoTransaction, repoItem, coupons, initialCoupons, validation)
	validation         = validator.New()
)

func TestInserTransactionService(t *testing.T) {
	data := &http_request.TransactionRequest{
		CustomerId: "bie7",
		CouponCode: "",
		PurchaseItems: []*http_request.TransactionItemsRequest{
			{
				ProductId: "p1",
				Quantity:  3,
			},
		},
	}

	result, err := transactionService.AddTransaction(ctx, data)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestGetTransaction(t *testing.T) {
	transactions, err := transactionService.GetTransaction(ctx, 1)
	assert.NoError(t, err)
	assert.NotEmpty(t, transactions)
	for i, transaction := range transactions {
		fmt.Println(i, "=>", transaction)
		for _, item := range transaction.PurchaseItems {
			fmt.Println(item)
		}
	}
}

func TestFindTransaction(t *testing.T) {
	transaction, err := transactionService.FindTransaction(ctx, "transaction-EUOM1lKXBcyQWSFy")
	assert.NoError(t, err)
	assert.NotEmpty(t, transaction)
	fmt.Println(transaction)
}

func TestDeleteTransaction(t *testing.T) {
	transaction, err := transactionService.DeleteTransaction(ctx, "transaction-EUOM1lKXBcyQWSFy")
	assert.NoError(t, err)
	assert.Equal(t, "Success Delete Transaction", transaction)
}
