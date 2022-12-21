package helper

import (
	entity2 "assigment-final-project/domain/entity/coupons"
	entity "assigment-final-project/domain/entity/transactions"
	"assigment-final-project/internal/delivery/http_request"
	"time"
)

func TransactionRequestToEntity(request *http_request.TransactionRequest, totalPriceProduct, totalAfterDiscount float64, discount float32) *entity.Transaction {
	return entity.NewTransaction(&entity.DTOTransaction{
		TransactionId:      "transaction-" + RandomString(16),
		CustomerId:         request.CustomerId,
		CouponCode:         request.CouponCode,
		TotalPrice:         totalPriceProduct,
		Discount:           discount,
		TotalAfterDiscount: totalAfterDiscount,
		PurchaseDate:       time.Now(),
	})
}

func CouponsRequestToEntity(couponCode, customerId string, expireDate time.Time) *entity2.Coupons {
	return entity2.NewCoupons(&entity2.DTOCoupons{
		CouponCode: couponCode,
		ExpireDate: expireDate,
		CustomerId: customerId,
	})
}

func TransactionItemRequestToEntity(request *http_request.TransactionItemsRequest, transactionId string) *entity.TransactionItems {
	return entity.NewTransactionItems(&entity.DTOTransactionItems{
		TransactionId: transactionId,
		ProductId:     request.ProductId,
		Quantity:      request.Quantity,
	})
}
