CREATE TABLE pvz IF NOT EXISTS pvz (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    registration_date TIMESTAMP NOT NULL,
    city VARCHAR(50) NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань'))
);

CREATE TABLE receptions IF NOT EXISTS receptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pvz_id UUID REFERENCES pvz(id),
    start_time TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('in_progress', 'closed'))
);

CREATE TABLE products IF NOT EXISTS products(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reception_id UUID REFERENCES receptions(id),
    added_time TIMESTAMP NOT NULL,
    product_type VARCHAR(20) NOT NULL CHECK (product_type IN ('electronics', 'clothing', 'shoes'))
);

CREATE TABLE users IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    role VARCHAR(100) NOT NULL CHECK (role IN ('client', 'moderator')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);