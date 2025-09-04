package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"})

	HTTPDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_duration_seconds",
			Help: "HTTP request latencies in seconds",
		},
		[]string{"path", "method", "status"})

	HTTPRequestsInFlight = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Number of HTTP requests currently in flight",
		},
		[]string{"path", "method"})

	HTTPErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP errors",
		},
		[]string{"path", "method", "status"},
	)

	HTTPSuccessCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_success_total",
			Help: "Total number of HTTP successful responses",
		},
		[]string{"path", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(
		HTTPRequestsTotal,
		HTTPDuration,
		HTTPRequestsInFlight,
		HTTPErrorCount,
		HTTPSuccessCount,
	)
}

// normalizePath reduces cardinality by normalizing dynamic paths
func normalizePath(path string) string {
	// Handle common API patterns
	if strings.HasPrefix(path, "/api/v") {
		parts := strings.Split(path, "/")
		if len(parts) >= 4 {
			// Replace dynamic IDs with placeholder
			for i := 4; i < len(parts); i++ {
				// If it looks like an ID (numeric or UUID-like), replace it
				if isNumeric(parts[i]) || isUUID(parts[i]) {
					parts[i] = "{id}"
				}
			}
			return strings.Join(parts, "/")
		}
	}

	// For other paths, keep as is but could add more normalization rules
	return path
}

func isNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(s) > 0
}

func isUUID(s string) bool {
	return len(s) == 36 && strings.Count(s, "-") == 4
}

func Prometheus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/metrics" {
			next.ServeHTTP(w, r)
			return
		}

		start := time.Now()
		path := normalizePath(r.URL.Path)
		method := r.Method

		// Increment in-flight requests (without status label)
		HTTPRequestsInFlight.WithLabelValues(path, method).Inc()
		defer HTTPRequestsInFlight.WithLabelValues(path, method).Dec()

		// Create a custom response writer to capture status code
		ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Process the request
		next.ServeHTTP(ww, r)

		// Record metrics after request completion
		status := strconv.Itoa(ww.statusCode)
		duration := time.Since(start).Seconds()

		// Record total requests
		HTTPRequestsTotal.WithLabelValues(path, method, status).Inc()

		// Record request duration
		HTTPDuration.WithLabelValues(path, method, status).Observe(duration)

		// Record success/error counts
		if ww.statusCode >= 400 {
			HTTPErrorCount.WithLabelValues(path, method, status).Inc()
		} else {
			HTTPSuccessCount.WithLabelValues(path, method, status).Inc()
		}
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	return rw.ResponseWriter.Write(data)
}
