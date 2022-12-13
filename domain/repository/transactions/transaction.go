package repository

import (
	entity "assigment-final-project/domain/entity/transactions"
	"context"
)

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) error
	FindTransaction(ctx context.Context, transactionId string) (*entity.Transaction, error)
	GetTransactions(ctx context.Context) ([]*entity.Transaction, error)
	DeleteTransaction(ctx context.Context, transactionId string) error
}
