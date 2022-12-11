package models

import "time"

type CouponsModel struct {
	Id         int       `dbq:"id"`
	CouponCode string    `dbq:"coupon_code"`
	IsUsed     bool      `dbq:"is_used"`
	ExpireDate time.Time `dbq:"expire_date"`
	CustomerId string    `dbq:"customer_id"`
}

func TableNameCoupons() string {
	return "coupons"
}
func FieldNameCoupons() []string {
	return []string{
		"id",
		"coupon_code",
		"is_used",
		"expire_date",
		"customer_id",
	}
}
