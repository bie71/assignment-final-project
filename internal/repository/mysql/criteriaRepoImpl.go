package repository

import (
	entity "assigment-final-project/domain/entity/criteria"
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

type CriteriaRepoImpl struct {
	db *sql.DB
}

func NewCriteriaRepoImpl(db *sql.DB) *CriteriaRepoImpl {
	return &CriteriaRepoImpl{db: db}
}

func (c *CriteriaRepoImpl) InsertCriteria(ctx context.Context, criteria *entity.Criteria) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := dbq.Struct(mapper.DomainCriteriaToCriteriaModel(criteria))
		stmt := dbq.INSERTStmt(models.TableNameCriteria(), models.FiledNameCriteria(), 1, dbq.MySQL)
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

func (c *CriteriaRepoImpl) GetCriteria(ctx context.Context) ([]*entity.Criteria, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s ", models.TableNameCriteria())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CriteriaModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ListModelCriteriaToListDomainCriteria(result.([]*models.CriteriaModel))
		return data, nil
	}
	return nil, errors.New("data empty")
}

func (c *CriteriaRepoImpl) UpdateCriteria(ctx context.Context, criteria *entity.Criteria) (*entity.Criteria, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`UPDATE %s SET name = ? WHERE id = ? `, models.TableNameCriteria())
		result, err := E(ctx, stmt, nil, criteria.CriteriaName())
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
			panic("Failed Update")
		} else {
			log.Println("Success Update ", affected)
		}

	})
	return criteria, errTx
}

func (c *CriteriaRepoImpl) DeleteCriteria(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ? `, models.TableNameCriteria())
		result, err := E(ctx, stmt, nil, id)
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
