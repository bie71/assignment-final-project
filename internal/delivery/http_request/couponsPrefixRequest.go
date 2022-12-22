package http_request

type CouponsPrefixRequest struct {
	PrefixName   string `json:"prefix_name,omitempty" validate:"required"`
	MinimumPrice int64  `json:"minimum_price,omitempty" validate:"required"`
	Discount     uint   `json:"discount,omitempty" validate:"required,gte=0,lte=100"`
	ExpireDate   string `json:"expire_date,omitempty" validate:"required"`
	Criteria     string `json:"criteria,omitempty"`
}
