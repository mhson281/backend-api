package middleware

import (
	"net/http"
	"strings"
	"context"

	"github.com/mhson281/backend-api/internal/auth"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// validate the token
		username, err := auth.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Pass the username to next handler (if needed)
		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
