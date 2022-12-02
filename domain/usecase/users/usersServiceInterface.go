package usecase

import (
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/internal/delivery/http_response"
	"context"
)

type UserService interface {
	AddUser(ctx context.Context, userRequest *http_request.UserRequest) (string, error)
	FindUser(ctx context.Context, UserLogin *http_request.UserLogin) (*http_response.UserResponse, error)
}
