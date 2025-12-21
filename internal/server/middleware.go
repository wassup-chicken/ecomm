package server

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"Method: %s | Path: %s | Duration: %v\n",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
