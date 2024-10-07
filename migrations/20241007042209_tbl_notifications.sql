-- +goose Up
-- +goose StatementBegin

-- Create table tbl_notifications
CREATE TABLE tbl_notifications (
    notification_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    message TEXT NOT NULL,
    status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER,
    updated_at TIMESTAMP NULL,
    updated_by INTEGER,
    deleted_at TIMESTAMPTZ,
    deleted_by INTEGER,
    FOREIGN KEY (user_id) REFERENCES tbl_users(user_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_notifications
DROP TABLE IF EXISTS tbl_notifications;

-- +goose StatementEnd
