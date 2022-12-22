package http_request

type TransactionRequest struct {
	CustomerId    string                     `json:"customer_id,omitempty" validate:"required"`
	CouponCode    string                     `json:"coupon_code,omitempty"`
	PurchaseItems []*TransactionItemsRequest `json:"purchase_items" validate:"required,dive"`
}

type TransactionItemsRequest struct {
	ProductId string `json:"product_id,omitempty" validate:"required"`
	Quantity  uint   `json:"quantity,omitempty" validate:"required,numeric"`
}
