package http_response

import (
	"assigment-final-project/domain/entity/users"
	"time"
)

type User struct {
	User any `json:"user"`
}

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

func ListDomainUserToListUserResponse(domain []*entity.Users) []*UserResponse {
	listResponse := make([]*UserResponse, 0)

	for _, user := range domain {
		result := DomainUsersToResponseUsers(user)
		listResponse = append(listResponse, result)
	}
	return listResponse
}
