package models

import "github.com/google/uuid"

type Wishlist struct {
	UserId    uuid.UUID `json:"wishlist_user_id" db:"wishlist_user_id"`
	ProductId uuid.UUID `json:"wishlist_product_id" db:"wishlist_product_id"`
}

type WishlistResponse struct {
	UserId   uuid.UUID `json:"wishlist_user_id"`
	Products []Product `json:"wishlist_products"`
}
