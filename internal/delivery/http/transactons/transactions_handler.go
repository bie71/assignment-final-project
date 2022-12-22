package handler

import "net/http"

type TransactionsHandler interface {
	AddTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactions(w http.ResponseWriter, r *http.Request)
	DeleteTransaction(w http.ResponseWriter, r *http.Request)
}
