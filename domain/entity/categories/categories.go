package entity

type Categories struct {
	categoryId string
	name       string
}

type DTOCategories struct {
	CategoryId string
	Name       string
}

func NewCategories(categories *DTOCategories) *Categories {
	return &Categories{
		categoryId: categories.CategoryId,
		name:       categories.Name,
	}
}

func (c *Categories) CategoryId() string {
	return c.categoryId
}

func (c *Categories) CategoryName() string {
	return c.name
}
