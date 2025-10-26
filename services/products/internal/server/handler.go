package server

import (
	"context"
	"log"

	api "buf.build/gen/go/wassup-chicken/common/protocolbuffers/go/api/v1"
)

func (s *ProductsServer) GetProduct(ctx context.Context, req *api.GetProductRequest) (*api.GetProductResponse, error) {
	// return grpc response
	log.Println("inv created")

	return &api.GetProductResponse{
		Id:      req.Id,
		Name:    "Hi!",
		Success: true,
	}, nil
}
