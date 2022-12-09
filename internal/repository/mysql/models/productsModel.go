package models

type ProductsModel struct {
	ProductId  string `dbq:"product_id"`
	Name       string `dbq:"name"`
	Price      uint   `dbq:"price"`
	CategoryId string `dbq:"category_id"`
	Stock      uint   `dbq:"stock"`
}

func TableNameProducts() string {
	return "products"
}
func FieldNameProducts() []string {
	return []string{
		"product_id",
		"name",
		"price",
		"category_id",
		"stock",
	}
}
