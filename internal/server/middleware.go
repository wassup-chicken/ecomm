package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

func (srv *JobServer) EnableCORS(next http.Handler) http.Handler {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://your-frontend-domain.com"}, // Specify allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                   // Specify allowed methods
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept"},                   // Specify allowed headers
		AllowCredentials: true,                                                                  // Allow cookies, HTTP auth, or client certs (only if precise origins are used)
		Debug:            true,                                                                  // Enable debugging for testing
	}

	return cors.New(corsOptions).Handler(next)
}

func (srv *JobServer) Logger(next http.Handler) http.Handler {
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

func (srv *JobServer) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement authentication
		// auth := r.Header.Get("Authorization")
		// token := strings.Split(auth, " ")
		// err := srv.Firebase.VerifyIDToken(r.Context(), token[1])
		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
