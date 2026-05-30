package middleware

import (
	"context"
	"instantanea/internal/handler"
	"instantanea/internal/service"
	"net/http"
	"strconv"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")

		if err != nil {
			http.Error(w, "JWT REQUIRED", http.StatusUnauthorized)
			return
		}

		claims, err := service.ValidateToken(cookie.Value)

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		token, err := service.GenerateToken(claims.Subject, claims.UserRole)

		if err != nil {
			http.Error(w, "Error interno al refrescar el token", http.StatusInternalServerError)
			return
		}

		handler.SetCookie(w, "access_token", token)

		userId, _ := strconv.Atoi(claims.Subject)

		ctx := context.WithValue(r.Context(), "UserID", userId)
		ctx  = context.WithValue(ctx, "UserRole", claims.UserRole)
		r    = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}