CREATE TABLE order_items(
    order_items_id UUID REFERENCES orders(order_id) ON DELETE CASCADE,
    order_items_product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    order_items_price NUMERIC(16,2) NOT NULL,
    order_items_quantity INT,
    PRIMARY KEY (order_items_id, order_items_product_id)
);