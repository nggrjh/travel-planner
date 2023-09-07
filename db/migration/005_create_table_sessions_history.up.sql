CREATE TABLE IF NOT EXISTS sessions_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    session_id UUID NOT NULL,
    token UUID NOT NULL,
    refreshed_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    operation VARCHAR NOT NULL,
    modified_by VARCHAR NOT NULL
);
