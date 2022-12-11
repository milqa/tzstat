package metrics

import (
	"net/http"
	"strconv"
	"time"
)

// responseWriterInterceptor is a simple wrapper to incercept set data on a http.ResponseWriter.
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (m *HttpMetrics) Middleware(handlerID string, inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wi := &responseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}

		hid := handlerID
		if handlerID == "" {
			hid = r.URL.Path
		}

		// Start the timer and when finishing measure the duration.
		start := time.Now()
		defer func() {
			duration := time.Since(start).Seconds()

			code := strconv.Itoa(wi.statusCode)

			m.ResponseTime.WithLabelValues(hid, r.Method, code).Observe(duration)
			m.Requests.WithLabelValues(hid, r.Method, code).Inc()
		}()

		inner.ServeHTTP(wi, r)
	})
}
