package repository

import (
	entity "assigment-final-project/domain/entity/transactions"
	"context"
)

type TransactionItemsRepo interface {
	InsertItems(ctx context.Context, items []*entity.TransactionItems) error
	GetItems(ctx context.Context) ([]*entity.TransactionItems, error)
	DeleteTransactionItems(ctx context.Context, id int) error
}
