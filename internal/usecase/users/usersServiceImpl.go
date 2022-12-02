package usecase

import (
	"assigment-final-project/domain/entity/users"
	users_repository "assigment-final-project/domain/repository/users"
	users_interface "assigment-final-project/domain/usecase/users"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

type ServiceUsersImplement struct {
	UserRepo   users_repository.UsersRepoInterface
	Validation *validator.Validate
}

func NewServiceUsersImplement(userRepo users_repository.UsersRepoInterface, validation *validator.Validate) users_interface.UserService {
	return &ServiceUsersImplement{UserRepo: userRepo, Validation: validation}
}

func (s *ServiceUsersImplement) AddUser(ctx context.Context, userRequest *http_request.UserRequest) (string, error) {
	errValidation := s.Validation.Struct(userRequest)
	if errValidation != nil {
		return "", errValidation
	}

	randomString := helper.RandomString(16)
	hashPassword := helper.HashAndSalt([]byte(userRequest.Password))
	userId := `user-` + randomString
	time, err := time.Parse(time.RFC1123Z, time.Now().Format(time.RFC1123Z))

	dataUser, err := entity.NewUsers(&entity.DTOUsers{
		UserId:    userId,
		Name:      userRequest.Name,
		Username:  userRequest.Username,
		Password:  hashPassword,
		UserType:  userRequest.UserType,
		CreatedAt: time,
	})
	helper.PanicIfError(err)

	username, err := s.UserRepo.FindUserByUsername(ctx, userRequest.Username)
	if username != nil {
		return "", errors.New("username already registered")
	}
	err = s.UserRepo.InsertUser(ctx, dataUser)
	if err != nil {
		return "", err
	}
	return dataUser.GetUserId(), nil
}

func (s *ServiceUsersImplement) FindUser(ctx context.Context, UserLogin *http_request.UserLogin) (*http_response.UserResponse, error) {
	errValidation := s.Validation.Struct(UserLogin)
	if errValidation != nil {
		return nil, errValidation
	}

	user, err := s.UserRepo.FindUserByUsername(ctx, UserLogin.Username)
	if err != nil {
		return nil, errors.New("user not yet registered")
	}

	if !helper.ComparePassword(user.Password(), []byte(UserLogin.Password)) {
		return nil, errors.New("password not match")
	}
	return http_response.DomainUsersToResponseUsers(user), nil
}
