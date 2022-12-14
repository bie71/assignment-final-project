package http_response

import "time"

type TransactionResponse struct {
	TransactionId           string                      `json:"transaction_id,omitempty"`
	Customer                *CustomerResponse           `json:"customer,omitempty"`
	CouponCode              string                      `json:"coupon_code,omitempty"`
	PurchaseItems           []*TransactionItemsResponse `json:"purchase_items"`
	TotalPrice              uint64                      `json:"total_price,omitempty"`
	Discount                int                         `json:"discount,omitempty"`
	TotalPriceAfterDiscount uint64                      `json:"total_price_after_discount,omitempty"`
	PurchaseDate            time.Time                   `json:"purchase_date"`
}

type TransactionItemsResponse struct {
	ProductId   string `json:"product_id,omitempty"`
	NameProduct string `json:"name_product,omitempty"`
	UnitPrice   uint   `json:"unit_price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
}
