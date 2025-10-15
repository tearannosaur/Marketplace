CREATE TABLE wishlist(
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE ,
    product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    PRIMARY KEY(user_id,product_id)
);
