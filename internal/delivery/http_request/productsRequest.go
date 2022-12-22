package http_request

type ProductsRequest struct {
	Name       string `json:"name,omitempty" validate:"required"`
	Price      uint   `json:"price,omitempty" validate:"required,numeric"`
	CategoryId string `json:"category_id,omitempty" validate:"required"`
	Stock      uint   `json:"stock,omitempty" validate:"required,numeric"`
}
