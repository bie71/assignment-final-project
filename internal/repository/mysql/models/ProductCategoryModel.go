package models

type ProductCategoryModel struct {
	ProductId    string `dbq:"product_id"`
	Name         string `dbq:"product_name"`
	Price        uint   `dbq:"price"`
	Stock        uint   `dba:"stock"`
	CategoryId   string `dbq:"category_id"`
	CategoryName string `dbq:"name"`
}
