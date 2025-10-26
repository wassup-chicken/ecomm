package main

import (
	"fmt"
	"log"
	"net"

	"github.com/wassup-chicken/ecomm/services/products/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "buf.build/gen/go/wassup-chicken/common/grpc/go/api/v1/apiv1grpc"
)

func main() {

	log.Println("hello from products service")

	ps := server.NewProducts()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", "8081"))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	api.RegisterProductServiceServer(grpcServer, ps)

	grpcServer.Serve(lis)
}
