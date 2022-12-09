package usecase

import (
	entity "assigment-final-project/domain/entity/customers"
	repository "assigment-final-project/domain/repository/customers"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type CustomerServiceImpl struct {
	customerRepo repository.RepoCustomer
	validation   *validator.Validate
}

func NewCustomerServiceImpl(customerRepo repository.RepoCustomer, validation *validator.Validate) *CustomerServiceImpl {
	return &CustomerServiceImpl{customerRepo: customerRepo, validation: validation}
}

func (c *CustomerServiceImpl) AddCustomer(ctx context.Context, customerRequest *http_request.CustomerRequest) (string, error) {
	errValidation := c.validation.Struct(customerRequest)
	if errValidation != nil {
		return "", errValidation
	}
	customerId := `customer-` + helper.RandomString(16)
	time, err := time.Parse(time.RFC1123Z, time.Now().Format(time.RFC1123Z))
	helper.PanicIfError(err)
	dataCutomer, _ := entity.NewCustomer(&entity.DTOCustomers{
		CustomerId: customerId,
		Name:       customerRequest.Name,
		Contact:    customerRequest.Phone,
		CreatedAt:  time,
	})

	customer, _ := c.customerRepo.FindCustomerById(ctx, dataCutomer.CustomerId(), dataCutomer.Contact())
	if customer != nil {
		return "", errors.New("customer already registered")
	}
	err = c.customerRepo.InsertCustomer(ctx, dataCutomer)
	helper.PanicIfError(err)
	return dataCutomer.CustomerId(), nil

}

func (c *CustomerServiceImpl) FindCustomer(ctx context.Context, customerId, phone string) (*http_response.CustomerResponse, error) {
	if customerId == "" && phone == "" {
		return nil, errors.New("customerid or phone must be set")
	}

	dataCustomer, err := c.customerRepo.FindCustomerById(ctx, customerId, phone)
	if err != nil {
		return nil, err
	}
	return http_response.DomainToCustomerResponse(dataCustomer), nil

}

func (c *CustomerServiceImpl) GetCustomers(ctx context.Context) ([]*http_response.CustomerResponse, error) {
	customers, err := c.customerRepo.GetCustomers(ctx)
	if err != nil {
		return nil, err
	}
	return http_response.ListDomainToListCustomerResponse(customers), nil
}

func (c *CustomerServiceImpl) DeleteCustomer(ctx context.Context, customerId, phone string) (string, error) {
	if customerId == "" && phone == "" {
		return "", errors.New("customerid or phone must be set")
	}
	_, err := c.customerRepo.FindCustomerById(ctx, customerId, phone)
	err = c.customerRepo.DeleteCustomerById(ctx, customerId, phone)
	if err != nil {
		return "", err
	}

	return "Success Delete Customer", nil
}
