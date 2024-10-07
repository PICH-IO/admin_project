-- +goose Up
-- +goose StatementBegin

-- Create table tbl_customers
CREATE TABLE tbl_customers (
    customer_id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_customers
DROP TABLE IF EXISTS tbl_customers;

-- +goose StatementEnd
