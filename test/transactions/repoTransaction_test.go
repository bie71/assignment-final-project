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
	repoTransaction.CreateTransaction(ctx, entity.NewTransaction(&entity.DTOTransaction{
		TransactionId: "1",
		CustomerId:    "bie7",
		CouponCode:    "",
		PurchaseDate:  time.Now(),
	}))
}
