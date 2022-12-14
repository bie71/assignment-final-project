package handler

import (
	usecase "assigment-final-project/domain/usecase/transactions"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type TransactionsHandlerImpl struct {
	transactionService usecase.TransactionService
}

func NewTransactionsHandlerImpl(transactionService usecase.TransactionService) *TransactionsHandlerImpl {
	return &TransactionsHandlerImpl{transactionService: transactionService}
}

func (t *TransactionsHandlerImpl) AddTransaction(w http.ResponseWriter, r *http.Request) {
	transactionRequest := &http_request.TransactionRequest{}
	helper.ReadFromRequestBody(r, transactionRequest)
	result, err := t.transactionService.AddTransaction(r.Context(), transactionRequest)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	delivery.ResponseDelivery(w, http.StatusCreated, result, nil)
}

func (t *TransactionsHandlerImpl) GetTransactions(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	transactionId := r.URL.Query().Get("id")
	if page == "" {
		page = "1"
	}
	p, err := strconv.Atoi(page)
	limit, _ := strconv.Atoi(os.Getenv("LIMIT"))
	helper.PanicIfError(err)

	if transactionId != "" {
		resultFind, err := t.transactionService.FindTransaction(r.Context(), transactionId)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, resultFind, nil)
		return
	}

	result, rows, err := t.transactionService.GetTransaction(r.Context(), p)
	if err != nil {
		fmt.Println(err)
	}

	delivery.ResponseDelivery(w, http.StatusOK, http_response.PaginationInfo(p, limit, rows, result), nil)
}

func (t *TransactionsHandlerImpl) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionId := r.URL.Query().Get("id")
	result, err := t.transactionService.DeleteTransaction(r.Context(), transactionId)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, nil, err)
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, result, nil)
}
