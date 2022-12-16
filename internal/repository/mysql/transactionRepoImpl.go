package repository

import (
	entity2 "assigment-final-project/domain/entity/products"
	entity "assigment-final-project/domain/entity/transactions"
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

type TransactionRepoImpl struct {
	db *sql.DB
}

func NewTransactionRepoImpl(db *sql.DB) *TransactionRepoImpl {
	return &TransactionRepoImpl{db: db}
}

func (t *TransactionRepoImpl) CreateTransaction(ctx context.Context, transaction *entity.Transaction) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, t.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := dbq.Struct(mapper.DomainTransactionToTransactionModel(transaction))
		stmt := dbq.INSERTStmt(models.TabelNameTransaction(), models.FieldNameTransaction(), 1, dbq.MySQL)
		result, errStore := E(ctx, stmt, nil, modelDbStruct)
		if errStore != nil {
			helper.PanicIfError(errStore)
			return
		}
		errCommit := txCommit()
		row, errCommit := result.RowsAffected()
		helper.PanicIfError(errCommit)
		log.Println("Succes Insert : ", row)
	})
	return errTx
}

func (t *TransactionRepoImpl) FindTransaction(ctx context.Context, transactionId string) (*entity.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE transaction_id = ?`, models.TabelNameTransaction())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.TransactionModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, t.db, stmt, opts, transactionId)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ModelToDomainTransaction(result.(*models.TransactionModel))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (t *TransactionRepoImpl) GetTransactions(ctx context.Context) ([]*entity.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, models.TabelNameTransaction())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.TransactionModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, t.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ListModelToListDomainTransaction(result.([]*models.TransactionModel))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (t *TransactionRepoImpl) DeleteTransaction(ctx context.Context, transactionId string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, t.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE transaction_id = ?`, models.TabelNameTransaction())
		result, err := E(ctx, stmt, nil, transactionId)
		if err != nil {
			helper.PanicIfError(err)
			return
		}

		errCommit := txCommit()
		helper.PanicIfError(errCommit)

		affected, err := result.RowsAffected()
		helper.PanicIfError(err)
		if affected == 0 {
			panic("Failed Delete")
		} else {
			log.Println("Success Delete", affected)
		}

	})
	return errTx
}

func (t *TransactionRepoImpl) GetProductJoinCategory(ctx context.Context, productId string) (*entity2.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT p.product_id, p.name as product_name, p.price, p.stock, c.category_id, c.name FROM %s p INNER JOIN %s c
                    ON p.category_id = c.category_id WHERE p.product_id = ?`, models.TableNameProducts(), models.TableNameCategories())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.ProductCategoryModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, t.db, stmt, opts, productId)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ModelToProductCategoryModel(result.(*models.ProductCategoryModel))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (t *TransactionRepoImpl) GetItemsProduct(ctx context.Context, transactionId string) ([]*entity.TransactionItemsProduct, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT ti.id, ti.transaction_id, p.product_id, p.name, p.price, ti.quantity FROM %s ti 
    INNER JOIN %s p ON ti.product_id = p.product_id WHERE ti.transaction_id = ?`,
		models.TableNameTransactionItems(), models.TableNameProducts())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.ItemsProductModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, t.db, stmt, opts, transactionId)
	helper.PanicIfError(err)
	if result != nil {
		return mapper.ListItemsProductToListItemsProductDomain(result.([]*models.ItemsProductModel)), nil
	}
	return nil, errors.New("data not found")
}
