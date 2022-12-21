package app

import (
	categories_handler "assigment-final-project/internal/delivery/http/categories"
	handler "assigment-final-project/internal/delivery/http/coupons"
	"assigment-final-project/internal/delivery/http/customers/customer_interface"
	products_handler "assigment-final-project/internal/delivery/http/products"
	"assigment-final-project/internal/delivery/http/users/users_interface"
	"assigment-final-project/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func Router(handlerUser users_interface.UserHandler, handlerCustomer customer_interface.CustomerHandler,
	handlerProduct products_handler.ProductsHandler, handlerCategory categories_handler.CategoryHandler,
	hanlderCoupons handler.CouponsHandler) *mux.Router {
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

	api.HandleFunc("/coupons/prefix", hanlderCoupons.AddCoupon).Methods(http.MethodPost)
	api.HandleFunc("/coupons/prefix", hanlderCoupons.GetCoupons).Methods(http.MethodGet)
	api.HandleFunc("/coupons/prefix", hanlderCoupons.UpdateAndDeleteCoupon).Queries("id", "{id}").Methods(http.MethodPut, http.MethodDelete)
	api.HandleFunc("/coupons", hanlderCoupons.GetCouponsCustomer).Queries("customerid", "{customerid}", "page", "{page}").Methods(http.MethodGet)
	return router
}
