package repository

import (
	entity "assigment-final-project/domain/entity/categories"
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

type CategoryRepoImpl struct {
	db *sql.DB
}

func NewCategoryRepoImpl(db *sql.DB) *CategoryRepoImpl {
	return &CategoryRepoImpl{db: db}
}

func (c *CategoryRepoImpl) InsertCategory(ctx context.Context, category *entity.Categories) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := dbq.Struct(mapper.DomainCategoryToCategoryModels(category))
		stmt := dbq.INSERTStmt(models.TableNameCategories(), models.FieldNameCategories(), 1, dbq.MySQL)
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

func (c *CategoryRepoImpl) FindCategory(ctx context.Context, categoryId string) (*entity.Categories, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE category_id = ?", models.TableNameCategories())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.CategoriesModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts, categoryId)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ModelCategoriesToDomainCategories(result.(*models.CategoriesModel))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (c *CategoryRepoImpl) GetCategories(ctx context.Context) ([]*entity.Categories, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s ", models.TableNameCategories())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CategoriesModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ListModelCategoriesToListDomainCategories(result.([]*models.CategoriesModel))
		return data, nil
	}
	return nil, errors.New("data empty")
}

func (c *CategoryRepoImpl) DeleteCategory(ctx context.Context, categoryId string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE category_id = ?`, models.TableNameCategories())
		result, err := E(ctx, stmt, nil, categoryId)
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
			log.Println("Success Delete ", affected)
		}

	})
	return errTx
}

func (c *CategoryRepoImpl) InsertListCategory(ctx context.Context, categories []*entity.Categories) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		listmodelDbStruct := mapper.DbqStructModelToListInterface(categories)
		stmt := dbq.INSERTStmt(models.TableNameCategories(), models.FieldNameCategories(), len(listmodelDbStruct), dbq.MySQL)
		result, errStore := E(ctx, stmt, nil, listmodelDbStruct)
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
