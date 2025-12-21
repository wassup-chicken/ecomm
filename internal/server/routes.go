package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (srv *JobServer) Routes() http.Handler {
	mux := chi.NewRouter()

	//middlewares
	mux.Use(Logger)
	//jobs
	mux.Get("/jobs", srv.GetJobs)
	mux.Get("/jobs/{id}", srv.GetJob)

	//users
	mux.Get("/users/{id}", srv.GetUser)
	// mux.Post("/users/login", srv.Login)
	mux.Post("/users/register", srv.Register)

	return mux
}
