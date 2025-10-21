CREATE TABLE products(
product_id UUID PRIMARY KEY,
product_price NUMERIC(16,2),
product_category TEXT REFERENCES category(category_name) ON DELETE CASCADE,
product_description TEXT
);