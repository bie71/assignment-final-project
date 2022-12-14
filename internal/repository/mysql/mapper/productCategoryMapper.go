package mapper

import (
	entity "assigment-final-project/domain/entity/categories"
	entity2 "assigment-final-project/domain/entity/products"
	"assigment-final-project/internal/repository/mysql/models"
)

func ModelToProductCategoryModel(model *models.ProductCategoryModel) *entity2.ProductCategory {
	category := entity.NewCategories(&entity.DTOCategories{
		CategoryId: model.CategoryId,
		Name:       model.CategoryName,
	})

	return &entity2.ProductCategory{
		ProductId:  model.ProductId,
		Name:       model.Name,
		Price:      model.Price,
		CategoryId: category,
		Stock:      model.Stock,
	}
}
