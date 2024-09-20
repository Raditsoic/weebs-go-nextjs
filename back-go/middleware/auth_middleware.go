package middleware

import "net/http"

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do some auth stuff
		next.ServeHTTP(w, r)
	})
}
