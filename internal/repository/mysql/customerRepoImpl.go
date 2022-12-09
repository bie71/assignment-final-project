package repository

import (
	entity "assigment-final-project/domain/entity/customers"
	"assigment-final-project/helper"
	"assigment-final-project/internal/repository/mysql/mapper"
	"assigment-final-project/internal/repository/mysql/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"log"
	"time"
)

type CustomerRepoImpl struct {
	db *sql.DB
}

func NewCustomerRepoImpl(db *sql.DB) *CustomerRepoImpl {
	return &CustomerRepoImpl{db: db}
}

func (c *CustomerRepoImpl) InsertCustomer(ctx context.Context, customer *entity.Customers) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := dbq.Struct(mapper.DomainToCustomersModel(customer))
		stmt := dbq.INSERTStmt(models.TableNameCustomer(), models.FieldNameCustomers(), 1, dbq.MySQL)
		result, errStore := E(ctx, stmt, nil, modelDbStruct)
		if errStore != nil {
			log.Println(errStore)
			return
		}
		errCommit := txCommit()
		row, errCommit := result.RowsAffected()
		helper.PanicIfError(errCommit)
		log.Println("Succes Insert : ", row)
	})
	return errTx
}

func (c *CustomerRepoImpl) FindCustomerById(ctx context.Context, customerId, phone string) (*entity.Customers, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE customer_id = ? OR contact = ?", models.TableNameCustomer())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.CustomerModels{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts, customerId, phone)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ModelsToDomainCustomers(result.(*models.CustomerModels))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (c *CustomerRepoImpl) GetCustomers(ctx context.Context) ([]*entity.Customers, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s ", models.TableNameCustomer())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CustomerModels{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ListModelToDomainListCustomer(result.([]*models.CustomerModels))
		return data, nil
	}
	return nil, errors.New("data empty")
}

func (c *CustomerRepoImpl) DeleteCustomerById(ctx context.Context, customerId, phone string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE customer_id = ? OR contact = ? `, models.TableNameCustomer())
		result, err := E(ctx, stmt, nil, customerId, phone)
		if err != nil {
			log.Println(err)
			return
		}

		errCommit := txCommit()
		helper.PanicIfError(errCommit)

		affected, err := result.RowsAffected()
		helper.PanicIfError(err)
		if affected == 0 {
			defer helper.RecoverPanic()
			panic("Failed Delete")
		} else {
			log.Println("Success Delete", affected)
		}

	})
	return errTx
}
