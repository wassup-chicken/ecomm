package server

import (
	"context"
	"log"

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

	product, err := s.ProductsStore.GetProduct(ctx, req.Id)

	if err != nil {
		log.Println("error getting product:", err)
		return nil, err
	}

	return &api.GetProductResponse{
		Id:        product.ProductID,
		Name:      product.ProductName,
		CreatedAt: timestamppb.New(product.CreatedAt), // convert to pb timestamp type
		UpdatedAt: timestamppb.New(product.UpdatedAt),
	}, nil
}
