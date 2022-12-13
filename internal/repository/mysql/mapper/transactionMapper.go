package mapper

import (
	entity "assigment-final-project/domain/entity/transactions"
	"assigment-final-project/internal/repository/mysql/models"
)

func DomainTransactionToTransactionModel(domain *entity.Transaction) *models.TransactionModel {
	return &models.TransactionModel{
		TransactionId:      domain.TransactionId(),
		CustomerId:         domain.CustomerId(),
		CouponCode:         domain.CouponCode(),
		TotalPrice:         domain.TotalPrice(),
		Discount:           domain.Discount(),
		TotalAfterDiscount: domain.TotalAfterDiscount(),
		PurchaseDate:       domain.PurchaseDate(),
	}
}

func ModelToDomainTransaction(model *models.TransactionModel) *entity.Transaction {
	return entity.NewTransaction(&entity.DTOTransaction{
		TransactionId:      model.TransactionId,
		CustomerId:         model.CustomerId,
		CouponCode:         model.CouponCode,
		TotalPrice:         model.TotalPrice,
		Discount:           model.Discount,
		TotalAfterDiscount: model.TotalAfterDiscount,
		PurchaseDate:       model.PurchaseDate,
	})
}

func ListModelToListDomainTransaction(models []*models.TransactionModel) []*entity.Transaction {
	listDomain := make([]*entity.Transaction, 0)

	for _, model := range models {
		transaction := ModelToDomainTransaction(model)
		listDomain = append(listDomain, transaction)
	}
	return listDomain
}
