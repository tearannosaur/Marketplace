CREATE TABLE wishlist(
    wishlist_user_id UUID REFERENCES users(user_id) ON DELETE CASCADE ,
    wishlist_product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    PRIMARY KEY(wishlist_user_id,wishlist_product_id)
);
