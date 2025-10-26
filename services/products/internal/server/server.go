package server

import (
	api "buf.build/gen/go/wassup-chicken/common/grpc/go/api/v1/apiv1grpc"
)

type ProductsServer struct {
	api.UnimplementedProductServiceServer
}

func NewProducts() *ProductsServer {
	return &ProductsServer{}
}
