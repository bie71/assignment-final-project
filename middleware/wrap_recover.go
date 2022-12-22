package middleware

import (
	"assigment-final-project/internal/delivery"
	"errors"
	"net/http"
)

func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				}
				delivery.ResponseDelivery(w, http.StatusInternalServerError, nil, err.Error())
			}
		}()
		h.ServeHTTP(w, r)
	})
}
