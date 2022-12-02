package handler

import "net/http"

type CategoryHandler interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	FindCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
}
