package main

import (
	"assigment-final-project/app"
	_ "assigment-final-project/helper"
	mysql_connection "assigment-final-project/internal/config/database/mysql"
	categories_handler "assigment-final-project/internal/delivery/http/categories"
	handler "assigment-final-project/internal/delivery/http/coupons"
	"assigment-final-project/internal/delivery/http/customers"
	products_handler "assigment-final-project/internal/delivery/http/products"
	users_handler "assigment-final-project/internal/delivery/http/users"
	repository "assigment-final-project/internal/repository/mysql"
	categories_service "assigment-final-project/internal/usecase/categories"
	usecase "assigment-final-project/internal/usecase/coupons"
	customers_service "assigment-final-project/internal/usecase/customers"
	products_service "assigment-final-project/internal/usecase/products"
	users_service "assigment-final-project/internal/usecase/users"
	"fmt"
	"github.com/go-playground/validator/v10"
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
	repoCoupons     = repository.NewCouponPrefixImpl(db)
	useCaseCoupons  = usecase.NewCouponsPrefixServiceImpl(repoCoupons, validate)
	couponsHandler  = handler.NewCouponHandlerImpl(useCaseCoupons)
)

func main() {

	router := app.Router(userHandler, customerHandler, productHandler, categoryHandler, couponsHandler)

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
