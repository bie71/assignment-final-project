package http_response

import (
	"assigment-final-project/domain/entity/users"
	"time"
)

type UserResponse struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}

func DomainUsersToResponseUsers(users *entity.Users) *UserResponse {
	return &UserResponse{
		UserId:    users.GetUserId(),
		Name:      users.GetName(),
		Username:  users.Username(),
		UserType:  users.UserType(),
		CreatedAt: users.CreatedAt(),
	}
}
