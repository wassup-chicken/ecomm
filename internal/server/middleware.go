package server

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/cors"
	"github.com/wassup-chicken/jobs/internal/util"
)

func (srv *Server) EnableCORS(next http.Handler) http.Handler {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://jobs-app-frontend.vercel.app"}, // Specify allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                       // Specify allowed methods
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept"},                       // Specify allowed headers
		AllowCredentials: true,                                                                      // Allow cookies, HTTP auth, or client certs (only if precise origins are used)
		Debug:            true,                                                                      // Enable debugging for testing
	}

	return cors.New(corsOptions).Handler(next)
}

func (srv *Server) Logger(next http.Handler) http.Handler {
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

func (srv *Server) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if len(auth) == 0 {
			util.ErrorJSON(w, errors.New("authorization token missing"), http.StatusUnauthorized)
			return
		}

		token := strings.Split(auth, " ")

		if len(token) != 2 {
			util.ErrorJSON(w, errors.New("authorization token missing"), http.StatusUnauthorized)
			return
		}

		err := srv.Firebase.VerifyIDToken(r.Context(), token[1])
		if err != nil {
			util.ErrorJSON(w, err, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
