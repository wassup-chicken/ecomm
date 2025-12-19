package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (srv *JobServer) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/jobs", srv.GetJobs)
	mux.Get("/jobs/{id}", srv.GetJob)

	return mux
}
