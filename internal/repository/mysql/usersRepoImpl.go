package repository

import (
	"assigment-final-project/domain/entity/users"
	usersInterface "assigment-final-project/domain/repository/users"
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

type UsersRepoImpl struct {
	db *sql.DB
}

func NewUsersRepoImpl(db *sql.DB) usersInterface.UsersRepoInterface {
	return &UsersRepoImpl{db: db}
}

func (u *UsersRepoImpl) InsertUser(ctx context.Context, dataUser *entity.Users) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	errTx := dbq.Tx(ctx, u.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		modelDbStruct := dbq.Struct(mapper.DomainUsersToModelsUsers(dataUser))
		stmt := dbq.INSERTStmt(models.TableNameUsers(), models.UsersFieldName(), 1, dbq.MySQL)
		result, errStore := E(ctx, stmt, nil, modelDbStruct)
		if errStore != nil {
			return
		}
		errCommit := txCommit()
		row, errCommit := result.RowsAffected()
		helper.PanicIfError(errCommit)
		log.Println("Succes Insert : ", row)
	})
	return errTx
}

func (u *UsersRepoImpl) GetUsers(ctx context.Context) ([]*entity.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, models.TableNameUsers())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.UsersModels{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, u.db, stmt, opts)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ToListDomainUser(result.([]*models.UsersModels))
		return data, nil
	}
	return nil, errors.New("data users is empty")
}

func (u *UsersRepoImpl) FindUserById(ctx context.Context, userId string) (*entity.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = ?`, models.TableNameUsers())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.UsersModels{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, u.db, stmt, opts, userId)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ModelsUsersToDomainUsers(result.(*models.UsersModels))
		return data, nil
	}
	return nil, errors.New("data users is not found")
}

func (u *UsersRepoImpl) FindUserByUsername(ctx context.Context, userName string) (*entity.Users, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE username = ?`, models.TableNameUsers())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.UsersModels{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, u.db, stmt, opts, userName)
	helper.PanicIfError(err)
	if result != nil {
		data := mapper.ModelsUsersToDomainUsers(result.(*models.UsersModels))
		return data, nil
	}
	return nil, errors.New("data users is not found")
}

func (u *UsersRepoImpl) UpdateById(ctx context.Context, dataUser *entity.Users, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, u.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`UPDATE %s 	SET name = ?, password = ?, user_type = ?, created_at = ? WHERE user_id = ? OR username = ?`, models.TableNameUsers())
		result, err := E(ctx, stmt, nil, dataUser.GetName(), dataUser.Password(), dataUser.UserType(), dataUser.CreatedAt(), userId, userId)
		if err != nil {
			return
		}

		errCommit := txCommit()
		helper.PanicIfError(errCommit)

		affected, err := result.RowsAffected()
		helper.PanicIfError(err)
		if affected == 0 {
			defer helper.RecoverPanic()
			panic("Failed Update User")
		} else {
			log.Println("Success Updated user ", affected)
		}

	})
	return errTx
}

func (u *UsersRepoImpl) DeleteById(ctx context.Context, userId string) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	errTx := dbq.Tx(ctx, u.db, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		stmt := fmt.Sprintf(`DELETE FROM %s WHERE user_id = ? OR username = ?`, models.TableNameUsers())
		result, err := E(ctx, stmt, nil, userId, userId)
		if err != nil {
			return
		}

		errCommit := txCommit()
		helper.PanicIfError(errCommit)

		affected, err := result.RowsAffected()
		helper.PanicIfError(err)
		if affected == 0 {
			defer helper.RecoverPanic()
			panic("Failed Delete User")
		} else {
			log.Println("Success Delete User ", affected)
		}

	})
	return errTx
}
