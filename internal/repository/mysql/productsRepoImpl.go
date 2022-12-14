package repository

import (
	entity "assigment-final-project/domain/entity/products"
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

type ProductsRepoImpl struct {
	db *sql.DB
}

func NewProductsRepoImpl(db *sql.DB) *ProductsRepoImpl {
	return &ProductsRepoImpl{db: db}
}

func (p *ProductsRepoImpl) InsertProduct(ctx context.Context, product *entity.Products) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, p.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := dbq.Struct(mapper.DomainProductsToProductsModel(product))
		stmt := dbq.INSERTStmt(models.TableNameProducts(), models.FieldNameProducts(), 1, dbq.MySQL)
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

func (p *ProductsRepoImpl) FindProduct(ctx context.Context, productId string) (*entity.Products, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE product_id = ?", models.TableNameProducts())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.ProductsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, p.db, stmt, opts, productId)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ProductsModelToDomainProducts(result.(*models.ProductsModel))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (p *ProductsRepoImpl) GetProducts(ctx context.Context) ([]*entity.Products, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s ", models.TableNameProducts())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.ProductsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, p.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ListModelProductsToListDomainProducts(result.([]*models.ProductsModel))
		return data, nil
	}
	return nil, errors.New("data empty")
}

func (p *ProductsRepoImpl) UpdateProduct(ctx context.Context, product *entity.Products) (*entity.Products, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, p.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`UPDATE %s SET name = ?, price = ?, category_id = ?, stock = ? WHERE product_id = ? `, models.TableNameProducts())
		result, err := E(ctx, stmt, nil, product.NameProduct(), product.Price(), product.Category(), product.Stock(), product.ProductId())
		if err != nil {
			helper.PanicIfError(err)
			return
		}

		errCommit := txCommit()
		helper.PanicIfError(errCommit)

		affected, err := result.RowsAffected()
		helper.PanicIfError(err)
		if affected == 0 {
			panic("Failed Update")
		} else {
			log.Println("Success Update ", affected)
		}

	})
	return product, errTx
}

func (p *ProductsRepoImpl) DeleteProduct(ctx context.Context, productId string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, p.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE product_id = ? `, models.TableNameProducts())
		result, err := E(ctx, stmt, nil, productId)
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

func (p *ProductsRepoImpl) InsertProducts(ctx context.Context, products []*entity.Products) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, p.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := mapper.DbqListProductToListInterface(products)
		stmt := dbq.INSERTStmt(models.TableNameProducts(), models.FieldNameProducts(), len(modelDbStruct), dbq.MySQL)
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
