package handler

import (
	usecase "assigment-final-project/domain/usecase/products"
	"assigment-final-project/helper"
	"assigment-final-project/internal/delivery"
	"assigment-final-project/internal/delivery/http_request"
	"github.com/go-playground/validator/v10"
	"net/http"
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

	response, err := p.serviceProducts.AddProduct(r.Context(), requestProduct)
	if err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, err.Error())
			return
		}

		delivery.ResponseDelivery(w, http.StatusBadRequest, errors.Error())
		return
	}

	delivery.ResponseDelivery(w, http.StatusCreated, response)
}

func (p *ProductsHandlerImpl) GetsFindAndDeleteProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")

	if query == "" {
		products, err := p.serviceProducts.GetProducts(r.Context())
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, products)
		return
	}

	if r.Method == http.MethodDelete {
		response, err := p.serviceProducts.DeleteProductById(r.Context(), query)
		if err != nil {
			delivery.ResponseDelivery(w, http.StatusNotFound, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusOK, response)
		return
	}

	response, err := p.serviceProducts.FindProductById(r.Context(), query)
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusNotFound, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, response)
}

func (p *ProductsHandlerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	requestProduct := &http_request.ProductsRequest{}
	helper.ReadFromRequestBody(r, requestProduct)

	product, err := p.serviceProducts.UpdateProduct(r.Context(), requestProduct, query)
	if err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			delivery.ResponseDelivery(w, http.StatusInternalServerError, err.Error())
			return
		}
		delivery.ResponseDelivery(w, http.StatusBadRequest, errors.Error())
		return
	}

	delivery.ResponseDelivery(w, http.StatusOK, product)
}

func (p *ProductsHandlerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := p.serviceProducts.GetProducts(r.Context())
	if err != nil {
		delivery.ResponseDelivery(w, http.StatusInternalServerError, err.Error())
		return
	}
	delivery.ResponseDelivery(w, http.StatusOK, products)
}
