-- +goose Up
-- +goose StatementBegin

-- Create table to link roles with permissions
CREATE TABLE tbl_role_permissions (
    role_id INT REFERENCES tbl_roles(role_id),
    permission_id INT REFERENCES tbl_permissions(permission_id),
    PRIMARY KEY (role_id, permission_id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER
);

-- Assign permissions to Admin and Sub Admin
-- Admin
INSERT INTO tbl_role_permissions (role_id, permission_id) VALUES 
(1, 1),  -- Admin can manage sub-admins
(1, 2),  -- Admin can manage products
(1, 3);  -- Admin can view reports

-- Sub Admin
INSERT INTO tbl_role_permissions (role_id, permission_id) VALUES 
(2, 2);  -- Sub Admin can manage products

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_role_permissions
-- +goose StatementEnd