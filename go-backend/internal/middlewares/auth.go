package middlewares

import (
	"instantanea/internal/helpers"
	"net/http"
)


func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")

		if err != nil {
			http.Error(w, "JWT REQUIRED", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		_, err = helpers.ValidateToken(tokenString)

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}