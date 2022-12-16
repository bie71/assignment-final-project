package entity

import "time"

type CouponsPrefix struct {
	id           int
	prefixName   string
	minimumPrice int64
	discount     uint
	expireDate   time.Time
	criteria     string
	createdAt    time.Time
}

type DTOCouponsPrefix struct {
	Id           int
	PrefixName   string
	MinimumPrice int64
	Discount     uint
	ExpireDate   time.Time
	Criteria     string
	CreatedAt    time.Time
}

func NewCouponsPrefix(DTO *DTOCouponsPrefix) *CouponsPrefix {
	return &CouponsPrefix{
		id:           DTO.Id,
		prefixName:   DTO.PrefixName,
		minimumPrice: DTO.MinimumPrice,
		discount:     DTO.Discount,
		expireDate:   DTO.ExpireDate,
		criteria:     DTO.Criteria,
		createdAt:    DTO.CreatedAt,
	}
}

func (c *CouponsPrefix) Id() int {
	return c.id
}

func (c *CouponsPrefix) PrefixName() string {
	return c.prefixName
}

func (c *CouponsPrefix) MinimumPrice() int64 {
	return c.minimumPrice
}

func (c *CouponsPrefix) Discount() uint {
	return c.discount
}

func (c *CouponsPrefix) ExpireDate() time.Time {
	return c.expireDate
}

func (c *CouponsPrefix) Criteria() string {
	return c.criteria
}

func (c *CouponsPrefix) CreatedAt() time.Time {
	return c.createdAt
}
