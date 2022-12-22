package repository

import (
	"assigment-final-project/domain/entity/users"
	"context"
)

type UsersRepoInterface interface {
	InsertUser(ctx context.Context, dataUser *entity.Users) error
	InsertUsers(ctx context.Context, dataUsers []*entity.Users) error
	GetUsers(ctx context.Context, offsetNum, limitNum int) ([]*entity.Users, error)
	FindUserById(ctx context.Context, userId string) (*entity.Users, error)
	FindUserByUsername(ctx context.Context, userName string) (*entity.Users, error)
	UpdateById(ctx context.Context, dataUser *entity.Users, userId string) error
	DeleteById(ctx context.Context, userId string) error
}
