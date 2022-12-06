package main

import (
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	"assigment-final-project/internal/delivery/http/customers"
	"assigment-final-project/internal/delivery/http/customers/customer_interface"
	products_handler "assigment-final-project/internal/delivery/http/products"
	users_handler "assigment-final-project/internal/delivery/http/users"
	"assigment-final-project/internal/delivery/http/users/users_interface"
	repository "assigment-final-project/internal/repository/mysql"
	customers_service "assigment-final-project/internal/usecase/customers"
	products_service "assigment-final-project/internal/usecase/products"
	user_service "assigment-final-project/internal/usecase/users"
	"assigment-final-project/middleware"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	db              = mysql_connection.InitMysqlDB()
	validate        = validator.New()
	repoUser        = repository.NewUsersRepoImpl(db)
	useCaseUser     = user_service.NewServiceUsersImplement(repoUser, validate)
	userHandler     = users_handler.NewUserHandlerImpl(useCaseUser)
	repoCustomer    = repository.NewCustomerRepoImpl(db)
	useCaseCustomer = customers_service.NewCustomerServiceImpl(repoCustomer, validate)
	customerHandler = customers.NewCustomerHandlerImpl(useCaseCustomer)
	repoProduct     = repository.NewProductsRepoImpl(db)
	useCaseProduct  = products_service.NewProductsServiceImpl(repoProduct, validate)
	productHandler  = products_handler.NewProductsHandlerImpl(useCaseProduct)
)

func Router(handler users_interface.UserHandler, handlerCustomer customer_interface.CustomerHandler, handlerProducts products_handler.ProductsHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", handler.Logout).Methods(http.MethodPost)

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthHandler)
	api.HandleFunc("/customer", handlerCustomer.AddCustomer).Methods(http.MethodPost)
	api.HandleFunc("/customer", handlerCustomer.GetAndDeleteCustomer).Queries("id", "{id}").Methods(http.MethodGet, http.MethodDelete)
	api.HandleFunc("/customer", handlerCustomer.GetAndDeleteCustomer).Methods(http.MethodGet)

	api.HandleFunc("/products", handlerProducts.AddProduct).Methods(http.MethodPost)
	api.HandleFunc("/products", handlerProducts.GetsFindAndDeleteProduct).Queries("id", "{id}").Methods(http.MethodGet, http.MethodDelete)
	api.HandleFunc("/products", handlerProducts.UpdateProduct).Queries("id", "{id}").Methods(http.MethodPut)
	api.HandleFunc("/products", handlerProducts.GetProducts).Methods(http.MethodGet)
	return router
}

func main() {

	router := Router(userHandler, customerHandler, productHandler)

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
