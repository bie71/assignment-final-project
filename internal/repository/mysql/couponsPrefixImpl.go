package repository

import (
	entity "assigment-final-project/domain/entity/coupons"
	"assigment-final-project/helper"
	"assigment-final-project/internal/repository/mysql/mapper"
	"assigment-final-project/internal/repository/mysql/models"
	"context"
	"database/sql"
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
	//TODO implement me
	panic("implement me")
}

func (c *CouponPrefixImpl) UpdatePrefix(ctx context.Context, prefix *entity.CouponsPrefix) (*entity.CouponsPrefix, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CouponPrefixImpl) DeletePrefix(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
