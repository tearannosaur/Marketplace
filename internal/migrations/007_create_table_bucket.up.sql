CREATE TABLE bucket(
    bucket_user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    bucket_product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    bucket_quantity INT,
    bucket_amount NUMERIC(16,2),
    PRIMARY KEY (bucket_user_id, bucket_product_id)
);