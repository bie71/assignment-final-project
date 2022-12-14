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
	transactionService = usecase.NewTransactionServiceImpl(repoTransaction, repoCutomer, repoItem, coupons, initialCoupons, validation)
	validation         = validator.New()
)

func TestInserTransactionService(t *testing.T) {
	item := []*http_request.TransactionItemsRequest{
		{
			ProductId: "p4",
			Quantity:  1,
		},
	}

	data := &http_request.TransactionRequest{
		TransactionId: "2",
		CustomerId:    "bie7",
		CouponCode:    "prime-je5FKTPDCq4vHwFe",
		PurchaseItems: item,
	}

	result, err := transactionService.AddTransaction(ctx, data)
	log.Println("error", err)
	fmt.Println(result)
}
