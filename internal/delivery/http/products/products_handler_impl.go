package handler

import (
	usecase "assigment-final-project/domain/usecase/products"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type ProductsHandlerImpl struct {
	serviceProducts usecase.ProductsService
}

func NewProductsHandlerImpl(serviceProducts usecase.ProductsService) *ProductsHandlerImpl {
	return &ProductsHandlerImpl{serviceProducts: serviceProducts}
}

func (p *ProductsHandlerImpl) AddProduct(w http.ResponseWriter, r *http.Request) {
	requestProduct := &http_request.ProductsRequest{}
	helper.ReadFromRequestBody(r, requestProduct)

	result, err := p.serviceProducts.AddProduct(r.Context(), requestProduct)
	if err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
			return
		}

		delivery.ResponseDelivery(w, http.StatusBadRequest, nil, errors.Error())
		return
	}

	delivery.ResponseDelivery(w, http.StatusCreated, result, nil)
}

func (p *ProductsHandlerImpl) GetsFindAndDeleteProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pg, err := strconv.Atoi(page)
	helper.PanicIfError(err)

	if query == "" {
		products, err := p.serviceProducts.GetProducts(r.Context(), pg)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, products, nil)
		return
	}

	if r.Method == http.MethodDelete {
		result, err := p.serviceProducts.DeleteProductById(r.Context(), query)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, result, nil)
		return
	}

	result, err := p.serviceProducts.FindProductById(r.Context(), query)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, result, nil)
}

func (p *ProductsHandlerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	requestProduct := &http_request.ProductsRequest{}
	helper.ReadFromRequestBody(r, requestProduct)

	result, err := p.serviceProducts.UpdateProduct(r.Context(), requestProduct, query)
	if err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusBadRequest, nil, errors.Error())
		return
	}

	delivery.ResponseDelivery(w, http.StatusOK, result, nil)
}

func (p *ProductsHandlerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pg, err := strconv.Atoi(page)
	helper.PanicIfError(err)
	result, err := p.serviceProducts.GetProducts(r.Context(), pg)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, result, nil)
}
