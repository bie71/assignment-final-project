package main

import (
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	"assigment-final-project/internal/delivery/http/customers"
	"assigment-final-project/internal/delivery/http/customers/customer_interface"
	Handler_Users "assigment-final-project/internal/delivery/http/users"
	"assigment-final-project/internal/delivery/http/users/users_interface"
	repository "assigment-final-project/internal/repository/mysql"
	customer_service "assigment-final-project/internal/usecase/customers"
	user_service "assigment-final-project/internal/usecase/users"
	"assigment-final-project/middleware"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Router(handler users_interface.UserHandler, handlerCustomer customer_interface.CustomerHandler) *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthHandler)
	api.HandleFunc("/customer", handlerCustomer.AddCustomer).Methods(http.MethodPost)
	api.HandleFunc("/customer", handlerCustomer.GetCustomer).Queries("id", "{id}").Methods(http.MethodGet)
	api.HandleFunc("/customer", handlerCustomer.GetCustomer).Methods(http.MethodGet)
	api.HandleFunc("/customer/{id}", handlerCustomer.DeleteCustomer).Methods(http.MethodDelete)

	router.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", handler.Logout).Methods(http.MethodPost)
	return router
}

func main() {
	db := mysql_connection.InitMysqlDB()
	validate := validator.New()
	repoUser := repository.NewUsersRepoImpl(db)
	useCaseUser := user_service.NewServiceUsersImplement(repoUser, validate)
	userHandler := Handler_Users.NewUserHandlerImpl(useCaseUser)
	repoCustomer := repository.NewCustomerRepoImpl(db)
	useCaseCustomer := customer_service.NewCustomerServiceImpl(repoCustomer, validate)
	customerHandler := customers.NewCustomerHandlerImpl(useCaseCustomer)
	router := Router(userHandler, customerHandler)

	server := http.Server{
		Addr:         "localhost:1919",
		Handler:      router,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}
	fmt.Println("server running on localhost:1919")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
