package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Ej0416/go-note-app/internal/json"
	"github.com/Ej0416/go-note-app/internal/types"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserContextKey contextKey = "user"

func Auth(jwtSecret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				json.Write(w,http.StatusUnauthorized, types.APIResponse{
					Success: false,
					Error: "missing authorization header",
				})
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				json.Write(w,http.StatusUnauthorized, types.APIResponse{
					Success: false,
					Error: "invalid authorization header",
				})
				return
			}

			tokenStr := parts[1]

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				json.Write(w,http.StatusUnauthorized, types.APIResponse{
					Success: false,
					Error: "invalid token",
				})
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				json.Write(w,http.StatusUnauthorized, types.APIResponse{
					Success: false,
					Error: "invalid token claims",
				})
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}