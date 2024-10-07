-- +goose Up
-- +goose StatementBegin

-- Create table tbl_categories
CREATE TABLE tbl_categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL UNIQUE,
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

-- Drop table tbl_categories
DROP TABLE IF EXISTS tbl_categories;

-- +goose StatementEnd
