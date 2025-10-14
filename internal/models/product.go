package models

import "github.com/google/uuid"

type Product struct {
	ProductId   uuid.UUID `json:"product_id"`
	Price       float64   `json:"product_price"`
	Category    string    `json:"product_category"`
	Description string    `json:"product_description"`
}

type ProductResponce struct {
	Price       float64 `json:"product_price"`
	Category    string  `json:"product_category"`
	Description string  `json:"product_description"`
}
