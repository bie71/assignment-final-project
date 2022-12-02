package jwt

import (
	"assigment-final-project/internal/delivery/http_response"
	"errors"
	"time"
)

type PayloadToken struct {
	Username  string `json:"username"`
	UserType  string `json:"user_type"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiredAt int64  `json:"expired_at"`
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func NewPayload(user *http_response.UserResponse, duration time.Duration) *PayloadToken {
	payload := &PayloadToken{
		Username:  user.Username,
		UserType:  user.UserType,
		IssuedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(duration).Unix(),
	}
	return payload
}

func (payload *PayloadToken) Valid() error {
	if time.Now().Unix() > payload.ExpiredAt {
		return ErrExpiredToken
	}
	return nil
}
