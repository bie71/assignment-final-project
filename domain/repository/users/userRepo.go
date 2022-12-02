package repository

import (
	"assigment-final-project/domain/entity/users"
	"context"
)

type UsersRepoInterface interface {
	InsertUser(ctx context.Context, dataUser *entity.Users) error
	GetUsers(ctx context.Context) ([]*entity.Users, error)
	FindUserById(ctx context.Context, userId string) (*entity.Users, error)
	FindUserByUsername(ctx context.Context, userName string) (*entity.Users, error)
	UpdateById(ctx context.Context, dataUser *entity.Users, userId string) error
	DeleteById(ctx context.Context, userId string) error
}
