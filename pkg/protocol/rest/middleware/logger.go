package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/bungysheep/contact-management/pkg/logger"
	"go.uber.org/zap"
)

// AddLogger add logger middleware
func AddLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var scheme string
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}

		proto := r.Proto
		method := r.Method
		remoteAddr := r.RemoteAddr
		userAgent := r.UserAgent()
		url := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")

		// Log Http request
		logger.Log.Debug("Request started", zap.String("http-scheme", scheme), zap.String("http-proto", proto), zap.String("http-method", method), zap.String("remote-addr", remoteAddr), zap.String("user-agent", userAgent), zap.String("url", url))

		t1 := time.Now()

		next.ServeHTTP(w, r)

		logger.Log.Debug("Request completed", zap.String("http-scheme", scheme), zap.String("http-proto", proto), zap.String("http-method", method), zap.String("remote-addr", remoteAddr), zap.String("user-agent", userAgent), zap.String("url", url), zap.Float64("elapsed-ms", float64(time.Since(t1).Nanoseconds())/1000000.0))
	})
}
