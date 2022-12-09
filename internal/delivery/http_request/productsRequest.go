package http_request

type ProductsRequest struct {
	Name       string `json:"name" validate:"required"`
	Price      uint   `json:"price" validate:"required,numeric"`
	CategoryId string `json:"category_id" validate:"required"`
	Stock      uint   `json:"stock" validate:"required,numeric"`
}
