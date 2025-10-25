package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (srv *GatewayServer) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/hello", srv.Hello)

	mux.Get("/products", srv.GetProducts)

	return mux
}
