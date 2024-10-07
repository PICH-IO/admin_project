-- +goose Up
-- +goose StatementBegin

-- Create table tbl_products
CREATE TABLE tbl_products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    category_id INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER,
    FOREIGN KEY (category_id) REFERENCES tbl_categories(category_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_products
DROP TABLE IF EXISTS tbl_products;

-- +goose StatementEnd
