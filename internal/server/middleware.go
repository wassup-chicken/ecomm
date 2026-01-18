package server

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/cors"
	"github.com/wassup-chicken/jobs/internal/util"
)

func (srv *Server) EnableCORS(next http.Handler) http.Handler {
	// Get allowed origins from environment variable
	// Format: comma-separated list, e.g., "http://localhost:5173,https://app.example.com"
	allowedOrigins := []string{"http://localhost:5173"} // Default for local dev
	
	corsOrigins := os.Getenv("CORS_ORIGINS")
	if corsOrigins != "" {
		// Split comma-separated origins and use them
		origins := strings.Split(corsOrigins, ",")
		allowedOrigins = make([]string, 0, len(origins))
		for _, origin := range origins {
			trimmed := strings.TrimSpace(origin)
			if trimmed != "" {
				allowedOrigins = append(allowedOrigins, trimmed)
			}
		}
	}

	corsOptions := cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept"},
		AllowCredentials: true,
		Debug:            os.Getenv("DEBUG") == "true", // Only debug if explicitly enabled
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
		err := srv.Firebase.VerifyIDToken(r.Context(), token[1])
		if err != nil {
			util.ErrorJSON(w, err, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
