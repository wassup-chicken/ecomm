package server

import (
	"context"

	api "buf.build/gen/go/wassup-chicken/common/protocolbuffers/go/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ProductsServer) GetProducts(ctx context.Context, req *api.GetProductsRequest) (*api.GetProductsResponse, error) {
	//returns all products
	products := []*api.GetProductResponse{}

	for i := range 2 {
		products = append(products, &api.GetProductResponse{
			Id:   int32(i),
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

	product := s.ProductsStore.GetProduct(ctx, req.Id)

	return &api.GetProductResponse{
		Id:        product.ProductID,
		Name:      "The coolest product on the market",
		Success:   true,
		CreatedAt: timestamppb.New(product.CreatedAt), // convert to pb timestamp type
	}, nil
}
