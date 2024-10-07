-- +goose Up
-- +goose StatementBegin

-- Create function to log changes in tbl_users
CREATE OR REPLACE FUNCTION audit_tbl_users() RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        INSERT INTO tbl_audit_logs (table_name, operation, record_id, changes, changed_by)
        VALUES ('tbl_users', TG_OP, NEW.user_id, row_to_json(NEW), NEW.created_by);
        RETURN NEW;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO tbl_audit_logs (table_name, operation, record_id, changes, changed_by)
        VALUES ('tbl_users', TG_OP, OLD.user_id, jsonb_build_object('old', row_to_json(OLD), 'new', row_to_json(NEW)), NEW.updated_by);
        RETURN NEW;
    ELSIF (TG_OP = 'DELETE') THEN
        INSERT INTO tbl_audit_logs (table_name, operation, record_id, changes, changed_by)
        VALUES ('tbl_users', TG_OP, OLD.user_id, row_to_json(OLD), OLD.deleted_by);
        RETURN OLD;
    END IF;
    RETURN NULL; -- Result is ignored since this is an AFTER trigger
END;
$$ LANGUAGE plpgsql;

-- Create trigger for tbl_users to call audit function after any changes
CREATE TRIGGER audit_users_trigger
AFTER INSERT OR UPDATE OR DELETE ON tbl_users
FOR EACH ROW EXECUTE FUNCTION audit_tbl_users();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop trigger and function for tbl_users
DROP TRIGGER IF EXISTS audit_users_trigger ON tbl_users;
DROP FUNCTION IF EXISTS audit_tbl_users;

-- +goose StatementEnd
