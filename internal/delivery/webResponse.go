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
	Message      string  `json:"message"`
	Error        bool    `json:"error"`
	ErrorMessage any     `json:"error_message"`
	Data         *Result `json:"data"`
}

type Result struct {
	Result any `json:"result"`
}

func ResponseDelivery(w http.ResponseWriter, code int, data any, errMsg any) {
	var err = false
	if errMsg != nil {
		err = true
	}

	response := &WebResponse{
		Code: code,
		Status: &Status{
			Message:      http.StatusText(code),
			Error:        err,
			ErrorMessage: errMsg,
			Data:         &Result{Result: data},
		},
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	helper.WriteToResponseBody(w, response)
}
