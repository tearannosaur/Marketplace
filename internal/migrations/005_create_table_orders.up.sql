CREATE TABLE orders(
    order_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    order_amount NUMERIC(16,2),
    order_status TEXT NOT NULL,
    order_created TIMESTAMP NOT NULL,
    order_updated TIMESTAMP
);