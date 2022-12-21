package repository

import (
	entity2 "assigment-final-project/domain/entity/products"
	entity "assigment-final-project/domain/entity/transactions"
	"context"
)

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
	FindTransaction(ctx context.Context, transactionId string) (*entity.Transaction, error)
	GetTransactions(ctx context.Context, offsetNum, limitNum int) ([]*entity.Transaction, error)
	DeleteTransaction(ctx context.Context, transactionId string) error
	GetProductJoinCategory(ctx context.Context, productId string) (*entity2.ProductCategory, error)
	GetItemsProduct(ctx context.Context, transactionId string) ([]*entity.TransactionItemsProduct, error)
	GetTransactionCustomers(ctx context.Context, transactionId string) (*entity.TransactionCustomer, error)
}
