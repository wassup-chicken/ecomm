package server

import (
	"context"
	"fmt"

	api "buf.build/gen/go/wassup-chicken/common/protocolbuffers/go/api/v1"
)

func (s *ProductsServer) GetProducts(ctx context.Context, req *api.GetProductsRequest) (*api.GetProductsResponse, error) {
	//returns all products
	products := []*api.GetProductResponse{}

	for i := range 2 {
		products = append(products, &api.GetProductResponse{
			Id:   fmt.Sprintf("%d", i),
			Name: "yo",
		})
	}
	// get products from postgres
	s.ProductsStore.GetProducts(ctx)

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
