-- +goose Up
-- +goose StatementBegin

-- Create table tbl_inventories
CREATE TABLE tbl_inventories (
    inventory_id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    stock_quantity INTEGER NOT NULL,
    last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER,
    FOREIGN KEY (product_id) REFERENCES tbl_products(product_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_inventories
DROP TABLE IF EXISTS tbl_inventories;

-- +goose StatementEnd
