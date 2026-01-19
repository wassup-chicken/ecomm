package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/wassup-chicken/jobs/internal/server"
)

func main() {
	log.Println("server starting")

	// Load .env file if it exists (for local development)
	// Don't fail if it doesn't exist (production uses environment variables)
	err := godotenv.Load()
	if err != nil {
		log.Println("unable to find .env file, using environment variables")
	}

	srv, err := server.New()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Get PORT from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Elastic Beanstalk sets PORT as just a number (e.g., "5000")
	// Go's http.Server needs it in format ":5000" or "0.0.0.0:5000"
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	log.Printf("server listening on %s", port)

	server := &http.Server{
		Addr:    port,
		Handler: srv.Routes(),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(fmt.Errorf("server failed to start: %w", err))
	}
}
