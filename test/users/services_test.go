package users_test

import (
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/repository/mysql"
	usecase "assigment-final-project/internal/usecase/users"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"testing"
)

var (
	ctx         = context.Background()
	db          = mysql_connection.InitMysqlDB()
	validate    = validator.New()
	repoUser    = repository.NewUsersRepoImpl(db)
	serviceuser = usecase.NewServiceUsersImplement(repoUser, validate)
)

func TestRegisterUser(t *testing.T) {
	user, err := serviceuser.AddUser(ctx, &http_request.UserRequest{
		Name:     "Habibi",
		Username: "bie7",
		Password: "1234",
		UserType: "owner",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
}

func TestLoginUser(t *testing.T) {
	user, err := serviceuser.FindUser(ctx, &http_request.UserLogin{
		Username: "bie7",
		Password: "1234",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
}
