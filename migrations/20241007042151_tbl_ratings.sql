-- +goose Up
-- +goose StatementBegin

-- Create table tbl_ratings
CREATE TABLE tbl_ratings (
    rating_id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    rating INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 5),
    review TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER,
    FOREIGN KEY (product_id) REFERENCES tbl_products(product_id),
    FOREIGN KEY (customer_id) REFERENCES tbl_customers(customer_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_ratings
DROP TABLE IF EXISTS tbl_ratings;

-- +goose StatementEnd
