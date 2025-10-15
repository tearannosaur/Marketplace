CREATE TABLE products(
product_id UUID PRIMARY KEY,
product_price NUMERIC(16,2),
category_id UUID REFERENCES category(category_id) ON DELETE CASCADE,
product_description TEXT
);