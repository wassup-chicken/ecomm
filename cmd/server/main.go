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

	server := http.Server{
		Addr:    ":8080",
		Handler: srv.Routes(),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
