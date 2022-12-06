package mapper

import (
	entity "assigment-final-project/domain/entity/products"
	"assigment-final-project/internal/repository/mysql/models"
)

func DomainProductsToProductsModel(products *entity.Products) *models.ProductsModel {
	return &models.ProductsModel{
		ProductId:         products.ProductId(),
		Name:              products.NameProduct(),
		Price:             products.Price(),
		CategoryId:        products.Category(),
		Stock:             products.Stock(),
		ProductsCondition: products.Condition(),
	}
}

func ProductsModelToDomainProducts(model *models.ProductsModel) *entity.Products {
	return entity.NewProducts(&entity.DTOProducts{
		ProductId:         model.ProductId,
		Name:              model.Name,
		Price:             model.Price,
		CategoryId:        model.CategoryId,
		Stock:             model.Stock,
		ProductsCondition: model.ProductsCondition,
	})
}

func ListModelProductsToListDomainProducts(list []*models.ProductsModel) []*entity.Products {
	listDomain := make([]*entity.Products, 0)

	for _, L := range list {
		product := ProductsModelToDomainProducts(L)
		listDomain = append(listDomain, product)
	}
	return listDomain
}
