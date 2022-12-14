package usecase

import (
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type TransactionService interface {
	AddTransaction(ctx context.Context, transactionRequest *http_request.TransactionRequest) (string, error)
	FindTransaction(ctx context.Context, transactionId string) (*http_response.TransactionResponse, error)
	GetTransaction(ctx context.Context) ([]*http_response.TransactionResponse, error)
	DeleteTransaction(ctx context.Context, transactionId string) (string, error)
}
