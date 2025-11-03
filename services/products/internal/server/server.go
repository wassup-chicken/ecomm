package server

import (
	api "buf.build/gen/go/wassup-chicken/common/grpc/go/api/v1/apiv1grpc"
	"github.com/wassup-chicken/ecomm/services/products/internal/store"
)

type ProductsServer struct {
	api.UnimplementedProductServiceServer
	ProductsStore store.ProductsStore
}

func NewProducts() *ProductsServer {
	//Initialize Clients if this connects to other apis

	//Initialize database
	productsStore := store.NewStore()

	return &ProductsServer{
		ProductsStore: productsStore,
	}
}
