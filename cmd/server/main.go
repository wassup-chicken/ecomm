package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/wassup-chicken/jobs/internal/server"
)

func main() {
	log.Println("server starting")

	err := godotenv.Load()

	if err != nil {
		log.Println("unable to find .env file")
		return
	}

	srv, err := server.New()

	if err != nil {
		log.Fatal(err)
		return
	}

	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: srv.Routes(),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
