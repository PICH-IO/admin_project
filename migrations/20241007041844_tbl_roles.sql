-- +goose Up
-- +goose StatementBegin

-- Create table tbl_roles
CREATE TABLE tbl_roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER
);

INSERT INTO tbl_roles (role_name) VALUES 
('Admin'), 
('Sub Admin'), 
('Customer');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_roles
DROP TABLE IF EXISTS tbl_roles;

-- +goose StatementEnd
