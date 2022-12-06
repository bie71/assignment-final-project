package http_request

type ProductsRequest struct {
	Name             string `json:"name" validate:"required"`
	Price            uint   `json:"price" validate:"required,numeric"`
	CategoryId       string `json:"category_id"`
	Stock            uint   `json:"stock" validate:"required,numeric"`
	ProductCondition string `json:"product_condition" validate:"required,eq=new|eq=second"`
}
