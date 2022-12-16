package entity

import "time"

type Transaction struct {
	transactionId      string
	customerId         string
	couponCode         string
	totalPrice         float64
	discount           float32
	totalAfterDiscount float64
	purchaseDate       time.Time
}

type DTOTransaction struct {
	TransactionId      string
	CustomerId         string
	CouponCode         string
	TotalPrice         float64
	Discount           float32
	TotalAfterDiscount float64
	PurchaseDate       time.Time
}

func NewTransaction(DTO *DTOTransaction) *Transaction {
	return &Transaction{
		transactionId:      DTO.TransactionId,
		customerId:         DTO.CustomerId,
		couponCode:         DTO.CouponCode,
		totalPrice:         DTO.TotalPrice,
		discount:           DTO.Discount,
		totalAfterDiscount: DTO.TotalAfterDiscount,
		purchaseDate:       DTO.PurchaseDate,
	}
}

func (t *Transaction) TransactionId() string {
	return t.transactionId
}

func (t *Transaction) CustomerId() string {
	return t.customerId
}

func (t *Transaction) CouponCode() string {
	return t.couponCode
}

func (t *Transaction) TotalPrice() float64 {
	return t.totalPrice
}

func (t *Transaction) Discount() float32 {
	return t.discount
}

func (t *Transaction) TotalAfterDiscount() float64 {
	return t.totalAfterDiscount
}

func (t *Transaction) PurchaseDate() time.Time {
	return t.purchaseDate
}
