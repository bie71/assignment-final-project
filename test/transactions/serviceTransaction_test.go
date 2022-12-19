package transactions_test

import (
	"assigment-final-project/internal/delivery/http_request"
	repoItems "assigment-final-project/internal/repository/mysql"
	usecase "assigment-final-project/internal/usecase/transactions"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"testing"
)

var (
	repoItem           = repoItems.NewTransactionItemsRepoImpl(db)
	repoCutomer        = repoItems.NewCustomerRepoImpl(db)
	initialCoupons     = repoItems.NewCouponPrefixImpl(db)
	coupons            = repoItems.NewCouponsRepoImpl(db)
	categoryRepo       = repoItems.NewCategoryRepoImpl(db)
	productRepo        = repoItems.NewProductsRepoImpl(db)
	transactionService = usecase.NewTransactionServiceImpl(repoTransaction, repoCutomer, repoItem,
		coupons, initialCoupons, productRepo, categoryRepo, validation)
	validation = validator.New()
)

func TestInserTransactionService(t *testing.T) {
	item := []*http_request.TransactionItemsRequest{
		{
			ProductId: "p2",
			Quantity:  4,
		},
	}

	data := &http_request.TransactionRequest{
		CustomerId:    "bie7",
		CouponCode:    "",
		PurchaseItems: item,
	}

	result, err := transactionService.AddTransaction(ctx, data)
	log.Println("error", err)
	fmt.Println(result)
}

func TestGetTransaction(t *testing.T) {
	transaction, err := transactionService.GetTransaction(ctx)
	fmt.Println(err)
	fmt.Println(transaction)
}
