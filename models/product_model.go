package models

import "time"

type Product struct {
	Id         string    `json:"id"` // UUID dari Supabase
	CreatedAt  time.Time `json:"created_at"`
	CategoryId string    `json:"category_id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	Category   Category  `json:"category"`
}

type ProductRequest struct {
	Id         string    `json:"id"` // UUID dari Supabase
	CreatedAt  time.Time `json:"created_at"`
	CategoryId string    `json:"category_id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
}
