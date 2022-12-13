package mapper

import (
	entity "assigment-final-project/domain/entity/categories"
	"assigment-final-project/internal/repository/mysql/models"
	"github.com/rocketlaunchr/dbq/v2"
)

func DomainCategoryToCategoryModels(category *entity.Categories) *models.CategoriesModel {
	return &models.CategoriesModel{
		CategoryId: category.CategoryId(),
		Name:       category.CategoryName(),
	}
}

func ModelCategoriesToDomainCategories(model *models.CategoriesModel) *entity.Categories {
	return entity.NewCategories(&entity.DTOCategories{
		CategoryId: model.CategoryId,
		Name:       model.Name,
	})
}

func ListModelCategoriesToListDomainCategories(list []*models.CategoriesModel) []*entity.Categories {
	listDomain := make([]*entity.Categories, 0)

	for _, model := range list {
		listDomain = append(listDomain, ModelCategoriesToDomainCategories(model))
	}
	return listDomain
}

func DbqStructModelToListInterface(listDomain []*entity.Categories) []interface{} {
	list := make([]interface{}, 0)

	for _, category := range listDomain {
		list = append(list, dbq.Struct(DomainCategoryToCategoryModels(category)))
	}
	return list
}
