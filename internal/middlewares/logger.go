package middlewares

import (
	"context"
	"net/http"

	logger_lib "github.com/s21platform/logger-lib"

	"github.com/s21platform/gateway-service/internal/config"
)

func LoggerMiddleware(next http.Handler, logger *logger_lib.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), config.KeyLogger, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
