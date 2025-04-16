CREATE TABLE pvz (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    registration_date TIMESTAMP NOT NULL,
    city VARCHAR(50) NOT NULL CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань'))
);

CREATE TABLE receptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pvz_id UUID REFERENCES pvz(id),
    start_time TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('in_progress', 'closed'))
);

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reception_id UUID REFERENCES receptions(id),
    added_time TIMESTAMP NOT NULL,
    product_type VARCHAR(20) NOT NULL CHECK (product_type IN ('electronics', 'clothing', 'shoes'))
);