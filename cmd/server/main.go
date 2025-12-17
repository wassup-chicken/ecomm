package main

import (
	"log"
	"net/http"

	"github.com/wassup-chicken/jobs/internal/handlers"
)

func main() {

	err := http.ListenAndServe(":8080", handlers.Routes())

	if err != nil {
		log.Fatal(err)
	}
}
