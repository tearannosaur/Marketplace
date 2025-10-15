CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    user_login TEXT UNIQUE NOT NULL,
    user_password TEXT NOT NULL,
    user_role TEXT NOT NULL,
    user_balance NUMERIC(16,2)
);

CREATE TABLE category(
    category_id UUID PRIMARY KEY ,
    category_name TEXT NOT NULL,
    category_description TEXT
);

CREATE TABLE products(
product_id UUID PRIMARY KEY,
product_price NUMERIC(16,2),
category_id UUID REFERENCES category(category_id) ON DELETE CASCADE,
product_description TEXT
);


CREATE TABLE wishlist(
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE ,
    product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    PRIMARY KEY(user_id,product_id)
);


CREATE TABLE orders(
    order_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    order_amount NUMERIC(16,2),
    order_status TEXT NOT NULL,
    order_created TIMESTAMP NOT NULL,
    order_updated TIMESTAMP
);

CREATE TABLE order_items(
    order_id UUID REFERENCES orders(order_id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    price NUMERIC(16,2) NOT NULL,
    quantity INT,
    PRIMARY KEY (order_id, product_id)
);



CREATE TABLE bucket(
    user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(product_id) ON DELETE CASCADE,
    quantity INT,
    amount NUMERIC(16,2),
    PRIMARY KEY (user_id, product_id)
);