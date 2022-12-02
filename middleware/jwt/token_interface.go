package jwt

import (
	"assigment-final-project/internal/delivery/http_response"
	"time"
)

type TokenJwt interface {
	CreateToken(user *http_response.UserResponse, duration time.Duration) (string, error)
	VerifyToken(token string) (*PayloadToken, error)
}
