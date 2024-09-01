CREATE TYPE addressable_type_enum AS ENUM ('user', 'store');

CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    label VARCHAR(255) NOT NULL,
    address_line_1 VARCHAR(255) NOT NULL,
    address_line_2 VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    postal_code VARCHAR(20),
    country VARCHAR(100) NOT NULL,
    addressable_id INT NOT NULL,
    addressable_type addressable_type_enum NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);