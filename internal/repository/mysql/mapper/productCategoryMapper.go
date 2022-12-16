package mapper

import (
	entity "assigment-final-project/domain/entity/categories"
	entity2 "assigment-final-project/domain/entity/products"
	entityItem "assigment-final-project/domain/entity/transactions"
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

func ItemsProdudtModelToDomain(model *models.ItemsProductModel) *entityItem.TransactionItemsProduct {
	return &entityItem.TransactionItemsProduct{
		Id:            model.Id,
		TransactionId: model.TransactionId,
		ProductId:     model.ProductId,
		Name:          model.Name,
		Price:         model.Price,
		Quantity:      model.Quantity,
	}
}

func ListItemsProductToListItemsProductDomain(list []*models.ItemsProductModel) []*entityItem.TransactionItemsProduct {
	listDomain := make([]*entityItem.TransactionItemsProduct, 0)
	for _, model := range list {
		listDomain = append(listDomain, ItemsProdudtModelToDomain(model))
	}
	return listDomain
}
