package repository

import (
	entity "assigment-final-project/domain/entity/coupons"
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

type CouponPrefixImpl struct {
	db *sql.DB
}

func NewCouponPrefixImpl(db *sql.DB) *CouponPrefixImpl {
	return &CouponPrefixImpl{db: db}
}

func (c *CouponPrefixImpl) InsertPrefix(ctx context.Context, prefix *entity.CouponsPrefix) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := dbq.Struct(mapper.DomainCounponsPrefixToModel(prefix))
		stmt := dbq.INSERTStmt(models.TableNameCouponsPrefix(), models.FieldNameCoupounsPrefix(), 1, dbq.MySQL)
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

func (c *CouponPrefixImpl) GetPrefixs(ctx context.Context, offsetNum, limitNum int) ([]*entity.CouponsPrefix, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s GROUP BY id LIMIT ?,?", models.TableNameCouponsPrefix())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CouponsPrefixModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts, offsetNum, limitNum)
	helper.PanicIfError(err)
	if result != nil {
		return mapper.ListModelToListDomainCouponsPrefix(result.([]*models.CouponsPrefixModel)), nil
	}
	return nil, errors.New("data empty")
}

func (c *CouponPrefixImpl) UpdatePrefix(ctx context.Context, prefix *entity.CouponsPrefix) (*entity.CouponsPrefix, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`UPDATE %s SET prefix_name = ?, minimum_price = ?, discount = ?, expire_date = ?, criteria = ? WHERE id = ?  `, models.TableNameCouponsPrefix())
		result, err := E(ctx, stmt, nil, prefix.PrefixName(), prefix.MinimumPrice(), prefix.Discount(), prefix.ExpireDate(), prefix.Criteria(), prefix.Id())
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
	return prefix, errTx
}

func (c *CouponPrefixImpl) DeletePrefix(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ? `, models.TableNameCouponsPrefix())
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

func (c *CouponPrefixImpl) InsertPrefixs(ctx context.Context, prefixs []*entity.CouponsPrefix) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := mapper.DbqStructCouponPrefixToListInterface(prefixs)
		stmt := dbq.INSERTStmt(models.TableNameCouponsPrefix(), models.FieldNameCoupounsPrefix(), len(modelDbStruct), dbq.MySQL)
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

func (c *CouponPrefixImpl) FindCouponPrefix(ctx context.Context, prefix, criteria string) (*entity.CouponsPrefix, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT ic.id, ic.prefix_name, ic.minimum_price, ic.discount, ic.expire_date, ic.created_at , c.name AS criteria"+
		" FROM %s ic INNER JOIN %s c ON ic.criteria = c.category_id WHERE ic.prefix_name = ? AND c.name = ? GROUP BY id", models.TableNameCouponsPrefix(), models.TableNameCategories())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.CouponsPrefixModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err := dbq.Q(ctx, c.db, stmt, opts, prefix, criteria)
	helper.PanicIfError(err)
	if result != nil {
		return mapper.ModelToDomainCounponsPrefix(result.(*models.CouponsPrefixModel)), nil
	}
	return nil, errors.New("data not found")
}

func (c *CouponPrefixImpl) GetPrefixMinimumPrice(ctx context.Context, price float64) ([]*entity.CouponsPrefix, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT DISTINCT  prefix_name , minimum_price , expire_date FROM %s WHERE minimum_price <= ? GROUP BY id", models.TableNameCouponsPrefix())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CouponsPrefixModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err := dbq.Q(ctx, c.db, stmt, opts, price)
	helper.PanicIfError(err)
	if result != nil {
		return mapper.ListModelToListDomainCouponsPrefix(result.([]*models.CouponsPrefixModel)), nil
	}
	return nil, errors.New("data not found")
}
