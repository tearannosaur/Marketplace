package models

import "github.com/google/uuid"

type Bucket struct {
	UserId    uuid.UUID `json:"user_id" db:"bucket_user_id"`
	ProductId uuid.UUID `json:"product_id" db:"bucket_product_id"`
	Quantity  float64   `json:"quantity" db:"bucket_quantity"`
	Amount    float64   `json:"price" db:"bucket_amount"`
}
