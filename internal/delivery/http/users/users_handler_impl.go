package Handler_Users

import (
	user_service "assigment-final-project/domain/usecase/users"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"assigment-final-project/middleware/jwt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	secretKey   = os.Getenv("SECRET_KEY")
	nameToken   = os.Getenv("NAME_TOKEN")
	duration, _ = strconv.Atoi(os.Getenv("DURATION"))
	expireTime  = time.Now().Add(time.Second * time.Duration(duration))
	NewJwt      = jwt.NewTokenJwtImpl(secretKey)
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

	result, err := u.UsersService.AddUser(r.Context(), userRequest)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusCreated, result, nil)
	return
}

func (u *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	userRequest := &http_request.UserLogin{}
	helper.ReadFromRequestBody(r, userRequest)

	result, err := u.UsersService.FindUser(r.Context(), userRequest)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		token, _ := NewJwt.CreateToken(result, time.Until(expireTime))
		http.SetCookie(w, &http.Cookie{
			Name:    nameToken,
			Value:   token,
			Expires: expireTime,
		})
	}()
	wg.Wait()

	delivery.ResponseDelivery(w, http.StatusOK, result, nil)
}

func (u *UserHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   nameToken,
		MaxAge: -1,
	})
	delivery.ResponseDelivery(w, http.StatusOK, "User has been logout", nil)
}

func (u *UserHandlerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := u.UsersService.GetUsers(r.Context())
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	delivery.ResponseDelivery(w, http.StatusOK, result, nil)
}
