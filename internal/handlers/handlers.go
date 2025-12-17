package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "hello world!")
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	r.Context().Value("id")
	log.Println(r.PathValue("id"))

	io.WriteString(w, "specific Job!!")
}

func GetCompany(w http.ResponseWriter, r *http.Request) {

	type user struct {
		ID   string
		Name string
	}

	u := user{
		ID:   "hi",
		Name: "seung",
	}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(u)

	if err != nil {
		log.Println(err)
	}
}
