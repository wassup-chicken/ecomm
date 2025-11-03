package store

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wassup-chicken/ecomm/services/products/models"
)

type ProductsStore interface {
	GetProduct(ctx context.Context, productId int32) *models.Products
	GetProducts(ctx context.Context) []*models.Products
	Close()
}

type productStore struct {
	url string
	//dbManager here postgres manager or connection
	dbPool *pgxpool.Pool
}

func NewStore() ProductsStore {
	dbpool, err := pgxpool.New(context.Background(), "postgresql://postgres@localhost:5432/products")
	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	//this shoudl be later // not in here
	// defer dbpool.Close()

	return &productStore{
		dbPool: dbpool,
	}
}

// Close closes the database connection pool
func (s *productStore) Close() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
}
