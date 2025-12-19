package main

import (
	"log"
	"net/http"

	"github.com/wassup-chicken/jobs/internal/server"
)

func main() {

	log.Println("server starting")

	srv, err := server.New()

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", srv.Routes())
}
