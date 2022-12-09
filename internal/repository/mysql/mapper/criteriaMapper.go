package mapper

import (
	entity "assigment-final-project/domain/entity/criteria"
	"assigment-final-project/internal/repository/mysql/models"
)

func DomainCriteriaToCriteriaModel(criteria *entity.Criteria) *models.CriteriaModel {
	return &models.CriteriaModel{
		Id:   criteria.Id(),
		Name: criteria.CriteriaName(),
	}
}

func CriteriaModelToDomainCriteria(criteria *models.CriteriaModel) *entity.Criteria {
	return entity.NewCriteria(&entity.DTOCriteria{
		Id:   criteria.Id,
		Name: criteria.Name,
	})
}

func ListModelCriteriaToListDomainCriteria(list []*models.CriteriaModel) []*entity.Criteria {
	listDomain := make([]*entity.Criteria, 0)
	for _, model := range list {
		criteria := CriteriaModelToDomainCriteria(model)
		listDomain = append(listDomain, criteria)
	}
	return listDomain
}
