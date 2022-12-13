package mapper

import (
	entity "assigment-final-project/domain/entity/transactions"
	"assigment-final-project/internal/repository/mysql/models"
	"github.com/rocketlaunchr/dbq/v2"
)

func DomainItemsToModelTransactionsItems(domain *entity.TransactionItems) *models.TransactionItemsModel {
	return &models.TransactionItemsModel{
		Id:            domain.Id(),
		TransactionId: domain.TransactionId(),
		ProductId:     domain.ProductId(),
		Quantity:      domain.Quantity(),
	}
}

func ModelItemsToDomainTransactionItems(model *models.TransactionItemsModel) *entity.TransactionItems {
	return entity.NewTransactionItems(&entity.DTOTransactionItems{
		Id:            model.Id,
		TransactionId: model.TransactionId,
		ProductId:     model.ProductId,
		Quantity:      model.Quantity,
	})
}

func ListModelItemsToListDomainTransactionItems(models []*models.TransactionItemsModel) []*entity.TransactionItems {
	listDomain := make([]*entity.TransactionItems, 0)
	for _, model := range models {
		listDomain = append(listDomain, ModelItemsToDomainTransactionItems(model))
	}
	return listDomain
}

func DbqStructModelTransactionItemsToListInterface(listModel []*entity.TransactionItems) []interface{} {
	list := make([]interface{}, 0)

	for _, model := range listModel {
		list = append(list, dbq.Struct(DomainItemsToModelTransactionsItems(model)))
	}
	return list
}
