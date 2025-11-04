package models

import "time"

type Products struct {
	ProductID   int32     `json:"product_id"`
	ProductName string    `json:"product_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
