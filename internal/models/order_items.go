package models

import "github.com/google/uuid"

type OrderItem struct {
	OrderId   uuid.UUID `json:"order_id"`
	ProductId uuid.UUID `json:"product_id"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
}
