package handler

import "net/http"

type ProductsHandler interface {
	AddProduct(w http.ResponseWriter, r *http.Request)
	GetsFindAndDeleteProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
}
