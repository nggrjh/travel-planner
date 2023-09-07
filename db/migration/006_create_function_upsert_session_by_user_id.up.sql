CREATE OR REPLACE FUNCTION upsert_session_by_user_id(
   user_id_in UUID,
   token_in UUID,
   refreshed_at_in TIMESTAMP,
   expires_at_in TIMESTAMP
)
   RETURNS VOID
   LANGUAGE plpgsql AS
$$

BEGIN
    INSERT INTO sessions (user_id, token, refreshed_at, expires_at, created_at)
        VALUES (user_id_in, token_in, refreshed_at_in, expires_at_in, NOW())
    ON CONFLICT (user_id)
    DO UPDATE SET
        token = token_in,
        refreshed_at = refreshed_at_in,
        expires_at = expires_at_in,
        last_updated_at = NOW();
END;

$$;
