package usecase

import (
	repository "assigment-final-project/domain/repository/customers"
	"assigment-final-project/helper"
	"assigment-final-project/helper/requestToEntity"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"os"
	"strconv"
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
	dataCutomer := requestToEntity.CustomerRequestToEntity(customerRequest, `customer-`+helper.RandomString(16))
	customer, _ := c.customerRepo.FindCustomerById(ctx, dataCutomer.CustomerId(), dataCutomer.Contact())
	if customer != nil {
		return "", errors.New("customer already registered")
	}
	err := c.customerRepo.InsertCustomer(ctx, dataCutomer)
	helper.PanicIfError(err)
	return dataCutomer.CustomerId(), nil

}

func (c *CustomerServiceImpl) FindCustomer(ctx context.Context, customerId, phone string) (*http_response.CustomerResponse, error) {
	if customerId == "" && phone == "" {
		return nil, errors.New("customerid or phone must be set")
	}

	dataCustomer, err := c.customerRepo.FindCustomerById(ctx, customerId, phone)
	if err != nil || dataCustomer == nil {
		return nil, errors.New("customer not found")
	}
	return http_response.DomainToCustomerResponse(dataCustomer), nil

}

func (c *CustomerServiceImpl) GetCustomers(ctx context.Context, page int) ([]*http_response.CustomerResponse, int, error) {
	var (
		limit, _ = strconv.Atoi(os.Getenv("LIMIT"))
		offset   = limit * (page - 1)
	)
	customers, err := c.customerRepo.GetCustomers(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	rows := helper.CountTotalRows(ctx, mysql_connection.InitMysqlDB(), "customers")
	return http_response.ListDomainToListCustomerResponse(customers), rows.TotalRows, nil
}

func (c *CustomerServiceImpl) DeleteCustomer(ctx context.Context, customerId, phone string) (string, error) {
	if customerId == "" && phone == "" {
		return "", errors.New("customerid or phone must be set")
	}
	result, err := c.customerRepo.FindCustomerById(ctx, customerId, phone)
	if err != nil || result == nil {
		return "", errors.New("customer not found")
	}
	err = c.customerRepo.DeleteCustomerById(ctx, customerId, phone)
	helper.PanicIfError(err)

	return "Success Delete Customer", nil
}
