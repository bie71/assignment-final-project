package models

import "time"

type TransactionModel struct {
	TransactionId      string    `dbq:"transaction_id"`
	CustomerId         string    `dbq:"customer_id"`
	CouponCode         string    `dbq:"coupon_code"`
	TotalPrice         uint64    `dbq:"total_price"`
	Discount           uint      `dbq:"discount"`
	TotalAfterDiscount uint64    `dbq:"total_price_after_discount"`
	PurchaseDate       time.Time `dbq:"purchase_date"`
}

func TabelNameTransaction() string {
	return "transaction"
}

func FieldNameTransaction() []string {
	return []string{
		"transaction_id",
		"customer_id",
		"coupon_code",
		"total_price",
		"discount",
		"total_price_after_discount",
		"purchase_date",
	}
}
