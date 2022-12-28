package usecase

import (
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type TransactionService interface {
	AddTransaction(ctx context.Context, transactionRequest *http_request.TransactionRequest) (*http_response.TransactionResult, error)
	FindTransaction(ctx context.Context, transactionId string) (*http_response.TransactionResponse, error)
	GetTransaction(ctx context.Context, page int) ([]*http_response.TransactionResponse, int, error)
	DeleteTransaction(ctx context.Context, transactionId string) (string, error)
}
