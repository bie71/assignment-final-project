package http_response

import entity "assigment-final-project/domain/entity/products"

type Product struct {
	Product any `json:"product"`
}

type ProductsResponse struct {
	ProductId  string `json:"product_id"`
	Name       string `json:"name"`
	Price      uint   `json:"price"`
	CategoryId string `json:"category_id"`
	Stock      uint   `json:"stock"`
}

func DomainProductsToProductsResponse(products *entity.Products) *ProductsResponse {
	return &ProductsResponse{
		ProductId:  products.ProductId(),
		Name:       products.NameProduct(),
		Price:      products.Price(),
		CategoryId: products.Category(),
		Stock:      products.Stock(),
	}
}

func ListDomainProductToResponseProducts(domain []*entity.Products) []*ProductsResponse {
	listResponse := make([]*ProductsResponse, 0)

	for _, D := range domain {
		response := DomainProductsToProductsResponse(D)
		listResponse = append(listResponse, response)
	}
	return listResponse
}
