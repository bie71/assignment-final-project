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

func (c *CouponPrefixImpl) GetPrefixs(ctx context.Context) ([]*entity.CouponsPrefix, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s ", models.TableNameCouponsPrefix())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CouponsPrefixModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ListModelToListDomainCouponsPrefix(result.([]*models.CouponsPrefixModel))
		return data, nil
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
	return prefix, errTx
}

func (c *CouponPrefixImpl) DeletePrefix(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ? `, models.TableNameCouponsPrefix())
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
