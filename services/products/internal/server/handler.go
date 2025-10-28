package server

import (
	"context"

	api "buf.build/gen/go/wassup-chicken/common/protocolbuffers/go/api/v1"
)

func (s *ProductsServer) GetProducts(ctx context.Context, req *api.GetProductsRequest) (*api.GetProductsResponse, error) {
	//returns all products
	products := []*api.GetProductResponse{}

	for i := 0; i < 2; i++ {
		products = append(products, &api.GetProductResponse{
			Id:   string(i),
			Name: "yo",
		})
	}
	return &api.GetProductsResponse{
		Products: products,
	}, nil
}

func (s *ProductsServer) GetProduct(ctx context.Context, req *api.GetProductRequest) (*api.GetProductResponse, error) {
	// return grpc response

	return &api.GetProductResponse{
		Id:      req.Id,
		Name:    "The coolest product on the market",
		Success: true,
	}, nil
}
