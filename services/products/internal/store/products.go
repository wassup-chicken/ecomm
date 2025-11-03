package store

import (
	"context"
	"log"

	"github.com/wassup-chicken/ecomm/services/products/models"
)

func (s *productStore) GetProduct(ctx context.Context, productId int32) *models.Products {
	product := &models.Products{}
	id := 22
	err := s.dbPool.QueryRow(ctx, "SELECT id, created_at from products where id = $1", id).Scan(&product.ProductID, &product.CreatedAt)

	if err != nil {
		log.Println("err running sql", err)
	}

	return &models.Products{
		ProductID: product.ProductID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func (s *productStore) GetProducts(ctx context.Context) []*models.Products {
	products := []*models.Products{}

	prod := &models.Products{
		ProductID: 1,
	}

	products = append(products, prod)
	return products
}

//TODO: Delete after testing

// func (s *productStore) InsertProduct(ctx context.Context) *models.Products {

// 	insertStatement := `INSERT INTO products`

// 	_, err := s.dbPool.Exec(ctx, insertStatement)

// 	if err != nil {
// 		log.Println(err)

// 	}

// 	var newId int

// 	log.Println("inserted!")
// }
