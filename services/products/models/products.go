package models

import "time"

type Products struct {
	ProductID int32     `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
