package models

import "time"

type CouponsPrefixModel struct {
	Id           int       `dbq:"id"`
	PrefixName   string    `dbq:"prefix_name"`
	MinimumPrice int64     `dbq:"minimum_price"`
	Discount     int       `dbq:"discount"`
	ExpireDate   time.Time `dbq:"expire_date"`
	Criteria     string    `dbq:"criteria"`
	CreatedAt    time.Time `dbq:"created_at"`
}

func TableNameCouponsPrefix() string {
	return "initial_coupons"
}
func FieldNameCoupounsPrefix() []string {
	return []string{
		"id",
		"prefix_name",
		"minimum_price",
		"discount",
		"expire_date",
		"criteria",
		"created_at",
	}
}
