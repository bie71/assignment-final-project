package mapper

import (
	entity "assigment-final-project/domain/entity/customers"
	"assigment-final-project/internal/repository/mysql/models"
)

func DomainToCustomersModel(customer *entity.Customers) *models.CustomerModels {
	return &models.CustomerModels{
		CustomerId: customer.CustomerId(),
		Name:       customer.Name(),
		Contact:    customer.Contact(),
		CreatedAt:  customer.CreatedAt(),
	}
}

func ModelsToDomainCustomers(customersModel *models.CustomerModels) *entity.Customers {
	return entity.CustomersFromDb(&entity.DTOCustomers{
		CustomerId: customersModel.CustomerId,
		Name:       customersModel.Name,
		Contact:    customersModel.Contact,
		CreatedAt:  customersModel.CreatedAt,
	})
}

func ListModelToDomainListCustomer(listModel []*models.CustomerModels) []*entity.Customers {
	listCustomer := make([]*entity.Customers, 0)

	for _, customerModels := range listModel {
		customers := ModelsToDomainCustomers(customerModels)
		listCustomer = append(listCustomer, customers)
	}
	return listCustomer
}
