package clients

import "log"

type ProductsClient interface {
	GetProducts()
}

type productClient struct {
	url string
}

func NewProducts() ProductsClient {
	log.Println("initialized product clients")

	return &productClient{
		url: "http://localhost:8081",
	}
}

func (c *productClient) GetProducts() {
	log.Println("hi!, get products called")

	//grpc connection to products service
}
