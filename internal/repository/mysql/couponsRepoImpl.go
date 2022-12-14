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

type CouponsRepoImpl struct {
	db *sql.DB
}

func NewCouponsRepoImpl(db *sql.DB) *CouponsRepoImpl {
	return &CouponsRepoImpl{db: db}
}

func (c *CouponsRepoImpl) FindCouponByCustomerIdAndCode(ctx context.Context, code, customerId string) (*entity.Coupons, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE coupon_code = ? AND customer_id = ?", models.TableNameCoupons())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.CouponsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts, code, customerId)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ModelCouponsToDomainCoupons(result.(*models.CouponsModel))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (c *CouponsRepoImpl) InsertCoupon(ctx context.Context, coupons *entity.Coupons) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := dbq.Struct(mapper.DomainCouponsToCouponsModel(coupons))
		stmt := dbq.INSERTStmt(models.TableNameCoupons(), models.FieldNameCoupons(), 1, dbq.MySQL)
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

func (c *CouponsRepoImpl) FindCoupon(ctx context.Context, code string, id int) (*entity.Coupons, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ? OR coupon_code = ?", models.TableNameCoupons())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.CouponsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts, id, code)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ModelCouponsToDomainCoupons(result.(*models.CouponsModel))
		return data, nil
	}
	return nil, errors.New("data not found")
}

func (c *CouponsRepoImpl) GetCoupons(ctx context.Context) ([]*entity.Coupons, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s ", models.TableNameCoupons())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CouponsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, c.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ListModelToListDomainCoupons(result.([]*models.CouponsModel))
		return data, nil
	}
	return nil, errors.New("data empty")
}

func (c *CouponsRepoImpl) DeleteCoupon(ctx context.Context, code string, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE id = ? OR coupon_code = ?`, models.TableNameCoupons())
		result, err := E(ctx, stmt, nil, id, code)
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
			log.Println("Success Delete ", affected)
		}

	})
	return errTx
}

func (c *CouponsRepoImpl) UpdateStatusCoupon(ctx context.Context, code, customerId string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, c.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`UPDATE %s SET is_used = ? WHERE coupon_code = ? AND customer_id = ? `, models.TableNameCoupons())
		result, err := E(ctx, stmt, nil, true, code, customerId)
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
	if errTx != nil {
		return false, errTx
	}
	return true, nil
}
