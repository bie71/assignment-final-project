package http_response

import entity "assigment-final-project/domain/entity/categories"

type Category struct {
	Category any `json:"category"`
}

type CategoryResponse struct {
	CategoryId string `json:"category_id"`
	Name       string `json:"name"`
}

func DomainToCategoryResponse(domain *entity.Categories) *CategoryResponse {
	return &CategoryResponse{
		CategoryId: domain.CategoryId(),
		Name:       domain.CategoryName(),
	}
}

func ListDomainToListCategory(listDomain []*entity.Categories) []*CategoryResponse {
	categories := make([]*CategoryResponse, 0)
	for _, category := range listDomain {
		response := DomainToCategoryResponse(category)
		categories = append(categories, response)
	}
	return categories
}
