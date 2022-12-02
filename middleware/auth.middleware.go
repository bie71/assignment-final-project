package middleware

import (
	"assigment-final-project/internal/delivery"
	"assigment-final-project/middleware/jwt"
	"log"
	"net/http"
	"os"
)

var (
	secretKey = os.Getenv("secret_key")
	NewJwt    = jwt.NewTokenJwtImpl(secretKey)
)

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("access_token")
		if err != nil {
			if err == http.ErrNoCookie {
				delivery.ResponseDelivery(w, http.StatusUnauthorized, nil)
				return
			}

			log.Println(err)
			delivery.ResponseDelivery(w, http.StatusBadRequest, nil)
			return
		}

		verifyToken, err := NewJwt.VerifyToken(token.Value)

		if err == nil {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", verifyToken.Username)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			delivery.ResponseDelivery(w, http.StatusForbidden, nil)
		}
	})
}
