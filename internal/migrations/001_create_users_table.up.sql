CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    user_login TEXT UNIQUE NOT NULL,
    user_password TEXT NOT NULL,
    user_role TEXT NOT NULL,
    user_balance NUMERIC(16,2)
);