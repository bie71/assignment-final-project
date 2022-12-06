package handler

import "net/http"

type CategoryHandler interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	FindAndDeleteCategory(w http.ResponseWriter, r *http.Request)
}
