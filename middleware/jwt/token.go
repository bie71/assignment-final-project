package jwt

import (
	"assigment-final-project/internal/delivery/http_response"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenJwtImpl struct {
	secretKey string
}

func NewTokenJwtImpl(secretKey string) *TokenJwtImpl {
	return &TokenJwtImpl{secretKey: secretKey}
}

func (t *TokenJwtImpl) CreateToken(user *http_response.UserResponse, duration time.Duration) (string, error) {
	payload := NewPayload(user, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(t.secretKey))
}

func (t *TokenJwtImpl) VerifyToken(token string) (*PayloadToken, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(t.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &PayloadToken{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*PayloadToken)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
