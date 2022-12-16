package models

type ProductCategoryModel struct {
	ProductId    string `dbq:"product_id"`
	Name         string `dbq:"product_name"`
	Price        uint   `dbq:"price"`
	Stock        uint   `dba:"stock"`
	CategoryId   string `dbq:"category_id"`
	CategoryName string `dbq:"name"`
}

type ItemsProductModel struct {
	Id            int    `dbq:"id"`
	TransactionId string `dbq:"transaction_id"`
	ProductId     string `dbq:"product_id"`
	Name          string `dbq:"name"`
	Price         uint   `dbq:"price"`
	Quantity      uint   `dbq:"quantity"`
}
