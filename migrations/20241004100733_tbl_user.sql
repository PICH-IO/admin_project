-- +goose Up
-- +goose StatementBegin

-- Create ENUM type for status
CREATE TYPE user_status AS ENUM ('active', 'inactive');

-- Create table tbl_users
CREATE TABLE tbl_users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    status user_status DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER
);

-- Insert initial admin user
INSERT INTO tbl_users (username, email, password, role_id, created_by)
VALUES ('admin', 'admintest123@gmail.com', '12345678', 1, 1);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_users and ENUM type
DROP TABLE IF EXISTS tbl_users;
DROP TYPE IF EXISTS user_status;

-- +goose StatementEnd
