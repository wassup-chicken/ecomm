package main

import (
	"log"

	"github.com/wassup-chicken/ecomm/services/products/internal/server"
)

func main() {

	log.Println("hello from products service")

	ps := server.NewProducts()

	log.Println(ps)

}
