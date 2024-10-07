-- +goose Up
-- +goose StatementBegin

-- Create table tbl_user_roles
CREATE TABLE tbl_user_roles (
    user_role_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER,
    FOREIGN KEY (user_id) REFERENCES tbl_users(user_id),
    FOREIGN KEY (role_id) REFERENCES tbl_roles(role_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_user_roles
DROP TABLE IF EXISTS tbl_user_roles;

-- +goose StatementEnd
