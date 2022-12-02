package jwt

import (
	"assigment-final-project/internal/delivery/http_response"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func TestNewTokenJwT(t *testing.T) {
	maker := NewTokenJwtImpl("rahasia")

	token, err := maker.CreateToken(&http_response.UserResponse{
		UserId:    "123",
		Name:      "habibi",
		Username:  "bie7",
		UserType:  "admin",
		CreatedAt: time.Now(),
	}, time.Second*10)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	verifyToken, err := maker.VerifyToken(token)

	assert.NotZero(t, verifyToken.Username)
	assert.Equal(t, "bie7", verifyToken.Username)
	assert.WithinDuration(t, time.Now(), time.Unix(verifyToken.IssuedAt, 0), 10*time.Second)
	assert.WithinDuration(t, time.Now(), time.Unix(verifyToken.ExpiredAt, 0), 10*time.Second)

	fmt.Println(token)
	fmt.Println(verifyToken)
}

func TestExpiredToken(t *testing.T) {
	maker := NewTokenJwtImpl("rahasia")

	token, err := maker.CreateToken(&http_response.UserResponse{
		UserId:    "123",
		Name:      "habibi",
		Username:  "bie7",
		UserType:  "admin",
		CreatedAt: time.Now(),
	}, -time.Second*10)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	verifyToken, err := maker.VerifyToken(token)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrExpiredToken.Error())
	assert.Nil(t, verifyToken)
	log.Println(err)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload := NewPayload(&http_response.UserResponse{
		UserId:    "bie",
		Name:      "habibi",
		Username:  "bie7",
		UserType:  "owner",
		CreatedAt: time.Now(),
	}, 10*time.Second)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker := NewTokenJwtImpl("rahasia")

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
	fmt.Println(err)
}
