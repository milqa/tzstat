package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	Http *HttpMetrics

	Handler http.Handler
}

func NewMetrics(app, version string) *Metrics {
	return &Metrics{
		Http:    NewHttpMetrics(app, version),
		Handler: promhttp.Handler(),
	}
}

type HttpMetrics struct {
	// ResponseTime is SummaryVec with labels handler, method, code.
	ResponseTime *prometheus.HistogramVec
}

func NewHttpMetrics(app, version string) *HttpMetrics {
	return &HttpMetrics{
		ResponseTime: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: app,
			Subsystem: "http",
			Name:      "request_duration_seconds",
			Help:      "Response len",
			ConstLabels: map[string]string{
				"version": version,
			},
		}, []string{"handler", "method", "code"}),
	}
}
