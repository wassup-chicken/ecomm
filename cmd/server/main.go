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

	// Support PORT variable (common in cloud platforms like Railway, Render, Fly.io)
	// Falls back to HOST if PORT is not set
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	
	var addr string
	if port != "" {
		addr = ":" + port
	} else if host != "" {
		addr = host
	} else {
		addr = ":8080" // Default fallback
	}

	log.Printf("Server starting on %s", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: srv.Routes(),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
