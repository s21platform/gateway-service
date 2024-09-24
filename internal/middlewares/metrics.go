package middlewares

import (
	"context"
	"github.com/s21platform/metrics-lib/pkg"
	"log"
	"net/http"
)

func MetricMiddleware(next http.Handler, metrics *pkg.Metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics.Increment("auth.login")
		log.Println("Increment for", r.URL.Path)
		ctx := context.WithValue(r.Context(), "metrics", metrics)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
