package metrics

import (
	"net/http"
	"time"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		RequestCount.WithLabelValues(r.Method, r.URL.Path).Inc()

		next.ServeHTTP(w, r)

		duration := time.Since(start).Seconds()
		RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}
