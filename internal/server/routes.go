package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (srv *JobServer) Routes() http.Handler {
	mux := chi.NewRouter()

	//middlewares
	mux.Use(srv.Authenticate)
	mux.Use(srv.Logger)

	//
	mux.Get("/", srv.Default)
	mux.Post("/upload", srv.Upload)

	//users
	mux.Get("/users/{id}", srv.GetUser)
	mux.Post("/users/register", srv.Register)
	mux.Post("/users/login", srv.Login)

	return mux
}
