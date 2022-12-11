package entity

import "time"

type Coupons struct {
	id         int
	couponCode string
	isUsed     bool
	expireDate time.Time
	customerId string
}

type DTOCoupons struct {
	Id         int
	CouponCode string
	IsUsed     bool
	ExpireDate time.Time
	CustomerId string
}

func NewCoupons(DTO *DTOCoupons) *Coupons {
	return &Coupons{
		id:         DTO.Id,
		couponCode: DTO.CouponCode,
		isUsed:     DTO.IsUsed,
		expireDate: DTO.ExpireDate,
		customerId: DTO.CustomerId,
	}
}

func (c *Coupons) Id() int {
	return c.id
}

func (c *Coupons) CouponCode() string {
	return c.couponCode
}

func (c *Coupons) IsUsed() bool {
	return c.isUsed
}

func (c *Coupons) ExpireDate() time.Time {
	return c.expireDate
}

func (c *Coupons) CustomerId() string {
	return c.customerId
}
