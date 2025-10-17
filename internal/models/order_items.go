package models

import "github.com/google/uuid"

type OrderItem struct {
	OrderId   uuid.UUID `json:"order_items_id" db:" order_items_id"`
	ProductId uuid.UUID `json:"product_id" db:"order_items_product_id"`
	Price     float64   `json:"price" db:"order_items_price"`
	Quantity  int       `json:"quantity" db:"order_items_quantity"`
}
