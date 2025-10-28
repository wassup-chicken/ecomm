package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (srv *GatewayServer) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/hello", srv.Hello)

	//get all products
	mux.Get("/products", srv.GetProducts)

	mux.Get("/products/{id}", srv.GetProduct)

	return mux
}
