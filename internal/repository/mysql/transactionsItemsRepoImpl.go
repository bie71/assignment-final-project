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

type TransactionItemsRepoImpl struct {
	db *sql.DB
}

func NewTransactionItemsRepoImpl(db *sql.DB) *TransactionItemsRepoImpl {
	return &TransactionItemsRepoImpl{db: db}
}

func (t *TransactionItemsRepoImpl) InsertItems(ctx context.Context, items []*entity.TransactionItems) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, t.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		listdbqStruct := mapper.DbqStructModelTransactionItemsToListInterface(items)
		stmt := dbq.INSERTStmt(models.TableNameTransactionItems(), models.FieldNameTransactionItems(), len(listdbqStruct), dbq.MySQL)
		result, errStore := E(ctx, stmt, nil, listdbqStruct)
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

func (t *TransactionItemsRepoImpl) GetItems(ctx context.Context) ([]*entity.TransactionItems, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, models.TableNameTransactionItems())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.TransactionItemsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, t.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ListModelItemsToListDomainTransactionItems(result.([]*models.TransactionItemsModel))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (t *TransactionItemsRepoImpl) DeleteTransactionItems(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, t.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, models.TableNameTransactionItems())
		result, err := E(ctx, stmt, nil, id)
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
