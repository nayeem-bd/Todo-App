package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nayeem-bd/Todo-App/internal/logger"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		logger.InfoWithFields("HTTP Request", logrus.Fields{
			"status":    ww.Status(),
			"method":    r.Method,
			"path":      r.URL.Path,
			"latency":   time.Since(start),
			"size":      ww.BytesWritten(),
			"client_ip": r.RemoteAddr,
		})
	})

}
