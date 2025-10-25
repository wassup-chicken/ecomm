package main

import (
	"log"
	"net/http"

	"github.com/wassup-chicken/ecomm/services/gateway/internal/server"
)

func main() {

	log.Println("Server starting on port 8080...")

	gs, err := server.NewGateway()

	srv := http.Server{
		Addr:    ":8080",
		Handler: gs.Routes(),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Println(err)

		log.Fatal("could not start http server for gateway")
	}
}
