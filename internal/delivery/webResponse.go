package delivery

import (
	"assigment-final-project/helper"
	"net/http"
)

type WebResponse struct {
	Code   int     `json:"code"`
	Status *Status `json:"status"`
}

type Status struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseDelivery(w http.ResponseWriter, code int, data interface{}) {
	response := &WebResponse{
		Code: code,
		Status: &Status{
			Message: http.StatusText(code),
			Data:    data,
		},
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	helper.WriteToResponseBody(w, response)
}
