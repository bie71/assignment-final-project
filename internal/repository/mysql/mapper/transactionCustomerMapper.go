package mapper

import (
	entity "assigment-final-project/domain/entity/transactions"
	"assigment-final-project/internal/repository/mysql/models"
)

func ModelToDomainTransactionCustomer(model *models.TransactionCustomer) *entity.TransactionCustomer {
	return &entity.TransactionCustomer{
		TransactionId: model.TransactionId,
		CustomerId:    model.CustomerId,
		Name:          model.Name,
		Contact:       model.Contact,
		CreatedAt:     model.CreatedAt,
	}
}
