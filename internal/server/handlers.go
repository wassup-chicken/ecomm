package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/wassup-chicken/jobs/internal/models"
	"github.com/wassup-chicken/jobs/internal/util"
)

func (srv *Server) Ping(w http.ResponseWriter, r *http.Request) {
	_ = util.WriteJSON(w, http.StatusOK, "PINGED@@@")

}

func (srv *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		util.ErrorJSON(w, err, http.StatusBadRequest)
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

	_ = util.WriteJSON(w, http.StatusOK, &user)
}

func (srv *Server) Register(w http.ResponseWriter, r *http.Request) {
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

func (srv *Server) Login(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")

	if len(auth) == 0 {
		util.ErrorJSON(w, errors.New("invalid auth"), http.StatusUnauthorized)
		return
	}

	token := strings.Split(auth, " ")[1]

	err := srv.Firebase.VerifyIDToken(r.Context(), token)

	if err != nil {
		util.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	util.WriteJSON(w, http.StatusOK, nil)
}

func (srv *Server) Upload(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	log.Printf("Upload request Content-Type: %s", contentType)

	file, header, err := r.FormFile("resume")

	if err != nil {
		util.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	url := r.FormValue("url")

	bs, err := io.ReadAll(file)

	if err != nil {
		util.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resumeb64 := base64.StdEncoding.EncodeToString([]byte(bs))

	res, err := srv.LLM.NewChatWithFile(r.Context(), url, resumeb64, header.Filename)

	if err != nil {
		util.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	util.WriteJSON(w, http.StatusOK, res)

}
