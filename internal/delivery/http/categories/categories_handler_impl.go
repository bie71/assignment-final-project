package handler

import (
	usecase "assigment-final-project/domain/usecase/categories"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"net/http"
	"strconv"
)

type CategoryHandlerImpl struct {
	categoryService usecase.CategoryService
}

func NewCategoryHandlerImpl(categoryService usecase.CategoryService) *CategoryHandlerImpl {
	return &CategoryHandlerImpl{categoryService: categoryService}
}

func (c *CategoryHandlerImpl) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var v interface{}
	helper.ReadFromRequestBody(r, &v)
	switch v := v.(type) {
	case []interface{}:
		// it's an array
		categories := make([]*http_request.CategoryRequest, 0)
		for _, data := range v {
			m := data.(map[string]interface{})
			s := m["name"].(string)
			categories = append(categories, &http_request.CategoryRequest{Name: s})
		}
		category, err := c.categoryService.AddCategories(r.Context(), categories)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusBadRequest, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusCreated, category, nil)
		return
	case map[string]interface{}:
		// it's an object
		s := v["name"].(string)
		category, err := c.categoryService.AddCategory(r.Context(), &http_request.CategoryRequest{Name: s})
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusBadRequest, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusCreated, category, nil)
	}

}

func (c *CategoryHandlerImpl) FindAndDeleteCategory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, err := strconv.Atoi(page)
	helper.PanicIfError(err)

	if r.Method == http.MethodDelete {
		data, err := c.categoryService.DeleteCategoryById(r.Context(), query)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, data, nil)
		return
	}

	if query == "" {
		data, err := c.categoryService.GetCategories(r.Context(), p)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, data, nil)
		return
	}

	data, err := c.categoryService.FindCategoryById(r.Context(), query)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, data, nil)
}

func (c *CategoryHandlerImpl) GetCategories(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, err := strconv.Atoi(page)
	helper.PanicIfError(err)
	data, err := c.categoryService.GetCategories(r.Context(), p)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, data, nil)
}
