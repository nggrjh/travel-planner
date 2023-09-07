CREATE OR REPLACE FUNCTION log_session_changes()
    RETURNS TRIGGER
    LANGUAGE plpgsql AS
$$

BEGIN
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' THEN
        INSERT INTO sessions_history (session_id, token, refreshed_at, expires_at, created_at, operation, modified_by)
            VALUES (NEW.id, NEW.token, NEW.refreshed_at, NEW.expires_at, NOW(), TG_OP, current_user);
    ELSIF TG_OP = 'DELETE' THEN
        INSERT INTO sessions_history (session_id, token, refreshed_at, expires_at, created_at, operation, modified_by)
            VALUES (OLD.id, OLD.token, OLD.refreshed_at, OLD.expires_at, NOW(), TG_OP, current_user);
    END IF;
    
    RETURN NEW;
END;

$$;

CREATE TRIGGER log_session
    AFTER INSERT OR UPDATE OR DELETE ON sessions
    FOR EACH ROW
        EXECUTE FUNCTION log_session_changes();
