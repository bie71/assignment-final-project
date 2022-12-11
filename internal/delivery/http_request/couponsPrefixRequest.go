package http_request

type CouponsPrefixRequest struct {
	PrefixName   string `json:"prefix_name" validate:"required"`
	MinimumPrice int64  `json:"minimum_price" validate:"required"`
	Discount     int    `json:"discount" validate:"required"`
	ExpireDate   string `json:"expire_date" validate:"required"`
	Criteria     string `json:"criteria"`
}
