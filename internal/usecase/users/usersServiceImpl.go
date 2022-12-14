package usecase

import (
	users_repository "assigment-final-project/domain/repository/users"
	users_interface "assigment-final-project/domain/usecase/users"
	"assigment-final-project/helper"
	"assigment-final-project/helper/requestToEntity"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"os"
	"strconv"
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

	dataUser, _ := requestToEntity.UserRequestToEntity(userRequest, `user-`+helper.RandomString(16))
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
	if err != nil || user == nil {
		return nil, errors.New("you are not registered")
	}

	if !helper.ComparePassword(user.Password(), []byte(UserLogin.Password)) {
		return nil, errors.New("invalid password")
	}
	return http_response.DomainUsersToResponseUsers(user), nil
}

func (s *ServiceUsersImplement) GetUsers(ctx context.Context, page int) ([]*http_response.UserResponse, int, error) {
	var (
		limit, _ = strconv.Atoi(os.Getenv("LIMIT"))
		offset   = limit * (page - 1)
	)

	users, err := s.UserRepo.GetUsers(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	rows := helper.CountTotalRows(ctx, mysql_connection.InitMysqlDB(), "users")
	return http_response.ListDomainUserToListUserResponse(users), rows.TotalRows, nil
}
