package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logging(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		duration := time.Since(start)
		logger.Info("HTTP request",
			zap.String("method", req.Method),
			zap.String("uri", req.RequestURI),
			zap.Duration("duration", duration),
		)
	})
}
