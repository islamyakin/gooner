package middleware

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func IPWhitelistMiddleware(allowedIPs []string, next http.Handler) mux.MiddlewareFunc {
	allowedIPSet := make(map[string]struct{})
	for _, ip := range allowedIPs {
		allowedIPSet[ip] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientIP := getClientIP(r)
			if _, allowed := allowedIPSet[clientIP]; !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getClientIP(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		return strings.Split(forwardedFor, ",")[0]
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}
