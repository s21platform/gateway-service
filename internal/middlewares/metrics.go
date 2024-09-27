package middlewares

import (
	"context"
	"github.com/s21platform/metrics-lib/pkg"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func MetricMiddleware(next http.Handler, metrics *pkg.Metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricString := strings.Trim(strings.Replace(r.URL.Path, "/", "_", -1), "_")
		metrics.Increment(metricString)
		t := time.Now()

		rec := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		ctx := context.WithValue(r.Context(), "metrics", metrics)
		next.ServeHTTP(rec, r.WithContext(ctx))

		duration := time.Since(t).Milliseconds()
		metrics.Increment(metricString + "." + strconv.Itoa(rec.status))
		metrics.Duration(duration, metricString)
	})
}
