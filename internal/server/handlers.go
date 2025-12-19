package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (srv *JobServer) GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := srv.jobStore.GetJobs(r.Context())

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	bs, err := json.Marshal(jobs)

	w.Write(bs)
}

func (srv *JobServer) GetJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Use the ID from the URL path
	job, err := srv.jobStore.GetJob(r.Context(), id)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	bs, err := json.Marshal(job)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {

}
