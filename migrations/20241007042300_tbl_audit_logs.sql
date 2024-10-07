-- +goose Up
-- +goose StatementBegin

-- Create table tbl_audit_logs
CREATE TABLE tbl_audit_logs (
    audit_id SERIAL PRIMARY KEY,
    table_name VARCHAR(100) NOT NULL,
    operation VARCHAR(10) NOT NULL,
    record_id INT NOT NULL,
    changes JSONB NOT NULL,
    changed_by INTEGER,
    changed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop table tbl_audit_logs
DROP TABLE IF EXISTS tbl_audit_logs;

-- +goose StatementEnd
