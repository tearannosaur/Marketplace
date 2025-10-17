package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId   uuid.UUID `json:"order_id" db:"order_id"`
	UserId    uuid.UUID `json:"order_user_id" db:"order_user_id"`
	Amount    float64   `json:"order_amount_id" db:"order_amount"`
	Status    string    `json:"order_status" db:"order_status"`
	CreatedAt time.Time `json:"order_created" db:"order_created"`
	UpdatedAt time.Time `json:"order_updated" db:"order_updated"`
}

type OrderResponce struct {
	OrderId   uuid.UUID `json:"order_id"`
	Amount    float64   `json:"order_amount_id"`
	Status    string    `json:"order_status"`
	CreatedAt time.Time `json:"order_created"`
	Items     []OrderItem
}
