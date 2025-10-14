package models

import "github.com/google/uuid"

type Bucket struct {
	UserId    uuid.UUID `json:"user_id"`
	ProductId uuid.UUID `json:"product_id"`
	Quantity  float64   `json:"quantity"`
	Amount    float64   `json:"price"`
}
