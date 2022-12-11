package delivery

import (
	"assigment-final-project/helper"
	"net/http"
)

type WebResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Status  *Status `json:"status"`
}

type Status struct {
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
		Code:    code,
		Message: http.StatusText(code),
		Status: &Status{
			Error:        err,
			ErrorMessage: errMsg,
			Data:         &Result{Result: data},
		},
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	helper.WriteToResponseBody(w, response)
}
