package models

import "github.com/google/uuid"

type Category struct {
	CategoryId   uuid.UUID `json:"category_id" db:"category_id"`
	CategoryName string    `json:"category_name" db:"category_name"`
	Description  string    `json:"category_description" db:"category_description"`
}
