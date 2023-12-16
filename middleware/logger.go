package middleware

import (
	"log"
	"net/http"
	"time"
	"strings"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		clientIP := getClientIPs(r)
		log.Printf("[%s] Started %s %s from %s", startTime.Format(time.RFC3339), r.Method, r.RequestURI, clientIP)
		next.ServeHTTP(w, r)
		duration := time.Since(startTime)

		log.Printf("[%s] Completed %s %s from %s in %v", startTime.Format(time.RFC3339), r.Method, r.RequestURI, clientIP, duration)
	})
}

func getClientIPs(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		return strings.TrimSpace(strings.Split(forwardedFor, ",")[0])
	}
	return strings.Split(strings.TrimSpace(r.RemoteAddr), ":")[0]
}
