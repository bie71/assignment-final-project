package handler

import (
	usecase "assigment-final-project/domain/usecase/categories"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"net/http"
)

type CategoryHandlerImpl struct {
	categoryService usecase.CategoryService
}

func (c *CategoryHandlerImpl) CreateCategory(w http.ResponseWriter, r *http.Request) {
	categoryRequest := &http_request.CategoryRequest{}
	helper.ReadFromRequestBody(r, categoryRequest)

	category, err := c.categoryService.AddCategory(r.Context(), categoryRequest)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusBadRequest, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusCreated, category)
}

func (c *CategoryHandlerImpl) FindAndDeleteCategory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")

	if r.Method == http.MethodDelete {
		data, err := c.categoryService.DeleteCategoryById(r.Context(), query)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusNotFound, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, data)
		return
	}

	if query == "" {
		data, err := c.categoryService.GetCategories(r.Context())
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, data)
		return
	}

	data, err := c.categoryService.FindCategoryById(r.Context(), query)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, data)
}
