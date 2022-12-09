package main

import (
	_ "assigment-final-project/db"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	categories_handler "assigment-final-project/internal/delivery/http/categories"
	"assigment-final-project/internal/delivery/http/customers"
	"assigment-final-project/internal/delivery/http/customers/customer_interface"
	products_handler "assigment-final-project/internal/delivery/http/products"
	users_handler "assigment-final-project/internal/delivery/http/users"
	"assigment-final-project/internal/delivery/http/users/users_interface"
	repository "assigment-final-project/internal/repository/mysql"
	categories_service "assigment-final-project/internal/usecase/categories"
	customers_service "assigment-final-project/internal/usecase/customers"
	products_service "assigment-final-project/internal/usecase/products"
	users_service "assigment-final-project/internal/usecase/users"
	"assigment-final-project/middleware"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	host            = os.Getenv("HOST")
	port            = os.Getenv("PORT")
	db              = mysql_connection.InitMysqlDB()
	validate        = validator.New()
	repoUser        = repository.NewUsersRepoImpl(db)
	useCaseUser     = users_service.NewServiceUsersImplement(repoUser, validate)
	userHandler     = users_handler.NewUserHandlerImpl(useCaseUser)
	repoCustomer    = repository.NewCustomerRepoImpl(db)
	useCaseCustomer = customers_service.NewCustomerServiceImpl(repoCustomer, validate)
	customerHandler = customers.NewCustomerHandlerImpl(useCaseCustomer)
	repoProduct     = repository.NewProductsRepoImpl(db)
	useCaseProduct  = products_service.NewProductsServiceImpl(repoProduct, validate)
	productHandler  = products_handler.NewProductsHandlerImpl(useCaseProduct)
	repoCategory    = repository.NewCategoryRepoImpl(db)
	useCaseCategory = categories_service.NewCategoryServiceImpl(repoCategory, validate)
	categoryHandler = categories_handler.NewCategoryHandlerImpl(useCaseCategory)
)

func Router(handlerUser users_interface.UserHandler, handlerCustomer customer_interface.CustomerHandler,
	handlerProduct products_handler.ProductsHandler, handlerCategory categories_handler.CategoryHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", handlerUser.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", handlerUser.Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", handlerUser.Logout).Methods(http.MethodPost)

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthHandler)
	api.HandleFunc("/users", handlerUser.GetUsers).Methods(http.MethodGet)

	api.HandleFunc("/customers", handlerCustomer.AddCustomer).Methods(http.MethodPost)
	api.HandleFunc("/customers", handlerCustomer.GetAndDeleteCustomer).Queries("id", "{id}").Methods(http.MethodGet, http.MethodDelete)
	api.HandleFunc("/customers", handlerCustomer.GetAndDeleteCustomer).Methods(http.MethodGet)

	api.HandleFunc("/categories", handlerCategory.CreateCategory).Methods(http.MethodPost)
	api.HandleFunc("/categories", handlerCategory.FindAndDeleteCategory).Queries("id", "{id}").Methods(http.MethodGet, http.MethodDelete)
	api.HandleFunc("/categories", handlerCategory.GetCategories).Methods(http.MethodGet)

	api.HandleFunc("/products", handlerProduct.AddProduct).Methods(http.MethodPost)
	api.HandleFunc("/products", handlerProduct.GetsFindAndDeleteProduct).Queries("id", "{id}").Methods(http.MethodGet, http.MethodDelete)
	api.HandleFunc("/products", handlerProduct.UpdateProduct).Queries("id", "{id}").Methods(http.MethodPut)
	api.HandleFunc("/products", handlerProduct.GetProducts).Methods(http.MethodGet)
	return router
}

func main() {

	router := Router(userHandler, customerHandler, productHandler, categoryHandler)

	server := http.Server{
		Addr:         host + ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}
	fmt.Printf("server running on %s:%s\n", host, port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
