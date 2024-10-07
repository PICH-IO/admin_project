-- +goose Up
-- +goose StatementBegin

-- Create function to log changes in tbl_customers
CREATE OR REPLACE FUNCTION audit_tbl_customers() RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
        INSERT INTO tbl_audit_logs (table_name, operation, record_id, changes, changed_by)
        VALUES ('tbl_customers', TG_OP, NEW.customer_id, row_to_json(NEW), NEW.created_by);
        RETURN NEW;
    ELSIF (TG_OP = 'UPDATE') THEN
        INSERT INTO tbl_audit_logs (table_name, operation, record_id, changes, changed_by)
        VALUES ('tbl_customers', TG_OP, OLD.customer_id, jsonb_build_object('old', row_to_json(OLD), 'new', row_to_json(NEW)), NEW.updated_by);
        RETURN NEW;
    ELSIF (TG_OP = 'DELETE') THEN
        INSERT INTO tbl_audit_logs (table_name, operation, record_id, changes, changed_by)
        VALUES ('tbl_customers', TG_OP, OLD.customer_id, row_to_json(OLD), OLD.deleted_by);
        RETURN OLD;
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Create trigger for tbl_customers to call audit function after any changes
CREATE TRIGGER audit_customers_trigger
AFTER INSERT OR UPDATE OR DELETE ON tbl_customers
FOR EACH ROW EXECUTE FUNCTION audit_tbl_customers();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop trigger and function for tbl_customers
DROP TRIGGER IF EXISTS audit_customers_trigger ON tbl_customers;
DROP FUNCTION IF EXISTS audit_tbl_customers;

-- +goose StatementEnd
