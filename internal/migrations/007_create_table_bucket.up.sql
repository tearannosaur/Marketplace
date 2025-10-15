CREATE TABLE bucket(
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    quantity INT,
    amount NUMERIC(16,2),
    PRIMARY KEY (user_id, product_id)
);