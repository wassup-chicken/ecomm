package store

import (
	"context"
	"log"

	"github.com/wassup-chicken/ecomm/services/products/models"
)

func (s *productStore) GetProduct(ctx context.Context, productId int32) (*models.Products, error) {
	product := &models.Products{}
	err := s.dbPool.QueryRow(ctx, "SELECT id, name, created_at, updated_at from products where id = $1", productId).Scan(&product.ProductID, &product.ProductName, &product.UpdatedAt, &product.CreatedAt)

	if err != nil {
		log.Println("err running sql", err)
		return nil, err
	}

	return &models.Products{
		ProductID:   product.ProductID,
		ProductName: product.ProductName,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
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
