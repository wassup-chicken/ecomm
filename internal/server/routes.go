package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (srv *Server) Routes() http.Handler {
	mux := chi.NewRouter()

	// Global middlewares (apply to all routes)
	mux.Use(srv.EnableCORS)
	mux.Use(srv.Logger)

	// Public routes (no authentication)
	mux.Get("/ping", srv.Ping)

	// Protected routes (require authentication)
	mux.Route("/", func(r chi.Router) {
		r.Use(srv.Authenticate)

		//resume load
		r.Post("/upload", srv.Upload)
	})

	return mux
}
