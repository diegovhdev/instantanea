package middlewares

import (
	"context"
	"instantanea/internal/helpers"
	"net/http"
	"strconv"
)


func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")

		if err != nil {
			http.Error(w, "JWT REQUIRED1", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		claims, err := helpers.ValidateToken(tokenString)

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userId, _ := strconv.Atoi(claims.Subject)

		ctx := context.WithValue(r.Context(), "UserID", userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}