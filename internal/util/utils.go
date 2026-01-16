package util

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	out, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	if err != nil {
		return err
	}
	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest

	// Use provided status code if given
	if len(status) > 0 {
		statusCode = status[0]
	}

	// Create error response structure
	errorResponse := map[string]string{
		"error": err.Error(),
	}

	// WriteJSON already writes to the writer, so we don't need to return the error
	WriteJSON(w, statusCode, errorResponse)
}
