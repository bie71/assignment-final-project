package http_response

import "time"

type TransactionResult struct {
	TransactionId string   `json:"transaction_id"`
	CouponCode    []string `json:"coupon_code"`
}

type TransactionResponse struct {
	TransactionId           string                      `json:"transaction_id"`
	Customer                *CustomerResponse           `json:"customer"`
	CouponCode              string                      `json:"coupon_code"`
	PurchaseItems           []*TransactionItemsResponse `json:"purchase_items"`
	TotalPrice              float64                     `json:"total_price"`
	Discount                float32                     `json:"discount"`
	TotalPriceAfterDiscount float64                     `json:"total_price_after_discount"`
	PurchaseDate            time.Time                   `json:"purchase_date"`
}

type TransactionItemsResponse struct {
	ProductId   string `json:"product_id,omitempty"`
	NameProduct string `json:"name_product,omitempty"`
	UnitPrice   uint   `json:"unit_price,omitempty"`
	Quantity    uint   `json:"quantity,omitempty"`
}
