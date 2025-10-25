package server

import (
	"github.com/wassup-chicken/ecomm/services/gateway/internal/clients"
)

type GatewayServer struct {
	productsClient clients.ProductsClient
}

func NewGateway() (*GatewayServer, error) {

	gs := GatewayServer{}

	// initialize clients

	gs.productsClient = clients.NewProducts()

	return &gs, nil
}
