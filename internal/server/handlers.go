package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (srv *JobServer) GetJobs(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)

	defer cancel()

	jobs, err := srv.jobStore.GetJobs(ctx)

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
	job, err := srv.jobStore.GetJob(ctx, intId)

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

	cook, _ := r.Cookie("This is Name")

	log.Println(cook)

	intId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		http.Error(w, "invalid user id!", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
	defer cancel()

	user, err := srv.jobStore.GetUser(ctx, intId)

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

type User struct {
	ID       uuid.UUID `json:"id"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}

var users = make(map[string]User)

func (srv *JobServer) Register(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	email := r.FormValue("email")
	password := r.FormValue("password")

	id := uuid.New()

	var user User
	if _, ok := users[email]; !ok {
		user = User{
			ID:       id,
			Email:    email,
			Password: password,
		}

		users[email] = user
	}

	user = users[email]

	bs, err := json.Marshal(user)

	if err != nil {
		log.Println(err)
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	w.Write(bs)
}

func (srv *JobServer) Login(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}
