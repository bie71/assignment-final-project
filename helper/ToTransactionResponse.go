package helper

import (
	entity "assigment-final-project/domain/entity/transactions"
	"assigment-final-project/internal/delivery/http_response"
)

func ToTransactionItemsResponse(item *entity.TransactionItemsProduct) *http_response.TransactionItemsResponse {
	return &http_response.TransactionItemsResponse{
		ProductId:   item.ProductId,
		NameProduct: item.Name,
		UnitPrice:   item.Price,
		Quantity:    item.Quantity,
	}
}
func ToTransactionsResponse(transaction *entity.Transaction, transactionCustomer *entity.TransactionCustomer, listItem []*http_response.TransactionItemsResponse) *http_response.TransactionResponse {
	dataCustomer := &http_response.CustomerResponse{
		CustomerId: transactionCustomer.CustomerId,
		Name:       transactionCustomer.Name,
		Phone:      transactionCustomer.Contact,
		CreatedAt:  transactionCustomer.CreatedAt,
	}

	return &http_response.TransactionResponse{
		TransactionId:           transaction.TransactionId(),
		Customer:                dataCustomer,
		CouponCode:              transaction.CouponCode(),
		PurchaseItems:           listItem,
		TotalPrice:              transaction.TotalPrice(),
		Discount:                transaction.Discount(),
		TotalPriceAfterDiscount: transaction.TotalAfterDiscount(),
		PurchaseDate:            transaction.PurchaseDate(),
	}
}
