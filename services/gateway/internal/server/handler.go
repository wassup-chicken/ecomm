package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (gs *GatewayServer) Hello(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusAccepted)

	w.Write([]byte("Yes, hi!"))
}

func (gs *GatewayServer) GetProducts(w http.ResponseWriter, r *http.Request) {

	gs.productsClient.GetProducts()

}

func (gs *GatewayServer) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")

	log.Println(productId)

	// gs.productsClient.GetProduct(productId)
}
