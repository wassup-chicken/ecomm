package server

import (
	"net/http"
)

func (gs *GatewayServer) Hello(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusAccepted)

	w.Write([]byte("Yes, hi!"))
}

func (gs *GatewayServer) GetProducts(w http.ResponseWriter, r *http.Request) {

	gs.productsClient.GetProducts()

}
