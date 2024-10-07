-- +goose Up
-- +goose StatementBegin

-- Create table tbl_permissions
CREATE TABLE tbl_permissions (
    permission_id SERIAL PRIMARY KEY,
    permission_name VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER
);

-- Insert initial permissions
INSERT INTO tbl_permissions (permission_name) VALUES 
('manage_sub_admins'),
('manage_products'),
('view_reports');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_permissions
DROP TABLE IF EXISTS tbl_permissions;

-- +goose StatementEnd
