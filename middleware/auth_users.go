package middleware

import (
	"assigment-final-project/internal/delivery"
	"log"
	"net/http"
	"strings"
)

func AuthUserHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _ := r.Cookie(nameToken)

		verifyToken, _ := NewJwt.VerifyToken(token.Value)

		if strings.EqualFold(verifyToken.UserType, "owner") {
			log.Printf("Authenticated user Owner %s\n", verifyToken.Username)
			next.ServeHTTP(w, r)
		} else {
			delivery.ResponseDelivery(w, http.StatusUnauthorized, nil, "You cannot access this resource")
		}
	})
}
