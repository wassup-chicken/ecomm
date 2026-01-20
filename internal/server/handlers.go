package server

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/wassup-chicken/jobs/internal/util"
)

func (srv *Server) Ping(w http.ResponseWriter, r *http.Request) {
	_ = util.WriteJSON(w, http.StatusOK, "pong")

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
		util.ErrorJSON(w, err)
		return
	}

	// res is already a JSON string from the LLM, parse it first to avoid double-encoding
	var parsedRes map[string]any
	if err := json.Unmarshal([]byte(res), &parsedRes); err != nil {
		// If parsing fails, return the string as-is (fallback)
		util.WriteJSON(w, http.StatusOK, res)
		return
	}

	util.WriteJSON(w, http.StatusOK, parsedRes)

}
