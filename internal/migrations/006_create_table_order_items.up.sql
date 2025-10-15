CREATE TABLE order_items(
    order_id UUID REFERENCES orders(order_id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    price NUMERIC(16,2) NOT NULL,
    quantity INT,
    PRIMARY KEY (order_id, product_id)
);