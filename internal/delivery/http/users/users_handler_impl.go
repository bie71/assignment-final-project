package Handler_Users

import (
	user_service "assigment-final-project/domain/usecase/users"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/middleware/jwt"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	secretKey  = os.Getenv("secret_key")
	expireTime = time.Now().Add(time.Hour * 24)
	NewJwt     = jwt.NewTokenJwtImpl(secretKey)
)

type UserHandlerImpl struct {
	UsersService user_service.UserService
}

func NewUserHandlerImpl(usersService user_service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{UsersService: usersService}
}

func (u *UserHandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	userRequest := &http_request.UserRequest{}
	helper.ReadFromRequestBody(r, userRequest)

	data, err := u.UsersService.AddUser(r.Context(), userRequest)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusBadRequest, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusCreated, data)
	return
}

func (u *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	userRequest := &http_request.UserLogin{}
	helper.ReadFromRequestBody(r, userRequest)

	data, err := u.UsersService.FindUser(r.Context(), userRequest)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusBadRequest, err.Error())
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		token, _ := NewJwt.CreateToken(data, time.Until(expireTime))
		http.SetCookie(w, &http.Cookie{
			Name:    "access_token",
			Value:   token,
			Expires: expireTime,
		})
	}()
	wg.Wait()

	delivery.ResponseDelivery(w, http.StatusOK, data)
	return
}

func (u *UserHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "access_token",
		MaxAge: -1,
	})
	delivery.ResponseDelivery(w, http.StatusOK, "User has been logout")
}
