package store

import "context"

// connections to database

type ProductsStore interface {
	GetProduct(ctx context.Context)
	GetProducts(ctx context.Context)
}

type dbClient struct {
	url string
}

func NewDatabse() ProductsStore {
	// create a connection pool to postgres

	return &dbClient{
		url: "postgresurl",
	}
}
