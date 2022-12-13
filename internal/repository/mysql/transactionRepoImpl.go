package repository

import (
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

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE transaction_id = ?`, models.TabelNameTransaction())
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
