package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/wassup-chicken/jobs/internal/models"
)

func (srv *JobServer) Ping(w http.ResponseWriter, r *http.Request) {
	sb, _ := json.Marshal([]byte("PINGED!!"))
	w.Write(sb)
}

func (srv *JobServer) GetJobs(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	jobs, err := srv.JobStore.GetJobs(ctx)

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
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)

	defer cancel()

	id := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("invalid id format:", err)
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	// Use the ID from the URL path
	job, err := srv.JobStore.GetJob(ctx, intId)

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

func (srv *JobServer) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		http.Error(w, "invalid user id!", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	user, err := srv.JobStore.GetUser(ctx, intId)

	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "no results found") {
			http.Error(w, "user not found!", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//probably need a session check here.

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bs, err := json.Marshal(user)

	if err != nil {
		log.Println(err)
		http.Error(w, "error forming json", http.StatusBadRequest)
		return
	}

	w.Write(bs)
}

func (srv *JobServer) Register(w http.ResponseWriter, r *http.Request) {
	log.Println(r)

	var user models.User

	bs, err := json.Marshal(user)

	if err != nil {
		log.Println(err)
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	w.Write(bs)
}

func (srv *JobServer) Login(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := srv.Firebase.VerifyIDToken(r.Context(), id)

	if err != nil {
		log.Println(err)

		return
	}

	log.Println(r)
}

func (srv *JobServer) Upload(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	log.Printf("Upload request Content-Type: %s", contentType)

	file, header, err := r.FormFile("resume")

	if err != nil {
		log.Printf("Failed to parse multipart form: %v", err)
		http.Error(w, "Failed to parse file. Ensure the request is sent as multipart/form-data with a 'resume' field. Error: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	url := r.FormValue("url")

	bs, err := io.ReadAll(file)

	resumeb64 := base64.StdEncoding.EncodeToString([]byte(bs))

	res, err := srv.LLM.NewChatWithFile(r.Context(), url, resumeb64, header.Filename)

	if err != nil {
		log.Println(err)
		io.WriteString(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, res)
}
