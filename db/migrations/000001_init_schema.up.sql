CREATE TYPE order_status_enum AS ENUM('INITIAL','PAID', 'ON THE WAY', 'COMPLETED', 'CANCELLED');

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    price FLOAT CHECK (price >= 0) NOT NULL,
    stock SMALLINT CHECK (stock >= 0) NOT NULL,
    category_id SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255) NOT NULL,
    isAdmin BOOLEAN NOT NULL,
    password VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS orderitems (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    product_id UUID NOT NULL,
    product_quantity SMALLINT CHECK (product_quantity >= 0) NOT NULL,
    product_price_agg FLOAT CHECK (product_price_agg >= 0) NOT NULL,
    order_id UUID NULL
);

CREATE TABLE IF NOT EXISTS orders(
    id UUID PRIMARY KEY,
    user_id BIGINT NOT NULL,
    status order_status_enum NOT NULL,
    payment_id UUID NOT NULL,
    total_price FLOAT CHECK (total_price >= 0) NOT NULL
);