CREATE TYPE product_detail_condition_enum AS ENUM ('baru', 'bekas');

CREATE TABLE product_details (
    id SERIAL PRIMARY KEY,
    condition product_detail_condition_enum NOT NULL,
    weight INT NOT NULL,
    description TEXT NOT NULL,
    price INT NOT NULL,
    quantity INT NOT NULL,
    minimum_purchase INT NOT NULL,
    photo VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);