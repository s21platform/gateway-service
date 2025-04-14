package adm

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/s21platform/gateway-service/internal/config"
)

func CheckJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Check jwt for:", r.URL.Path)

		// Получаем токен из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("no authorization header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Проверяем формат Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("invalid authorization header format")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		ctx := context.WithValue(r.Context(), config.KeyStaffUUID, tokenString)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
