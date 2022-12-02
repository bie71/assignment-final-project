package entity

import (
	"errors"
	"time"
)

type Users struct {
	userId    string
	name      string
	username  string
	password  string
	userType  string
	createdAt time.Time
}

type DTOUsers struct {
	UserId    string
	Name      string
	Username  string
	Password  string
	UserType  string
	CreatedAt time.Time
}

func NewUsers(DTOUser *DTOUsers) (*Users, error) {

	if DTOUser.UserId == "" {
		return nil, errors.New("userId required")
	}
	if DTOUser.Name == "" {
		return nil, errors.New("name required")
	}
	if DTOUser.Username == "" {
		return nil, errors.New("username required")
	}
	if DTOUser.UserType == "" {
		return nil, errors.New("usertype required")
	}

	return &Users{
		userId:    DTOUser.UserId,
		name:      DTOUser.Name,
		username:  DTOUser.Username,
		password:  DTOUser.Password,
		userType:  DTOUser.UserType,
		createdAt: DTOUser.CreatedAt,
	}, nil
}

func UserFromDb(users *DTOUsers) *Users {
	return &Users{
		userId:    users.UserId,
		name:      users.Name,
		username:  users.Username,
		password:  users.Password,
		userType:  users.UserType,
		createdAt: users.CreatedAt,
	}
}

func (u *Users) GetUserId() string {
	return u.userId
}
func (u *Users) GetName() string {
	return u.name
}

func (u *Users) Username() string {
	return u.username
}

func (u *Users) Password() string {
	return u.password
}

func (u *Users) UserType() string {
	return u.userType
}

func (u *Users) CreatedAt() time.Time {
	return u.createdAt
}
