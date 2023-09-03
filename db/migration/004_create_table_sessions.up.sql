CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    user_id UUID NOT NULL,
    token UUID NOT NULL,
    refreshed_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    last_updated_at TIMESTAMP,
    CONSTRAINT uidx_sessions_user_id UNIQUE (user_id)
);

CREATE INDEX idx_sessions_token ON sessions(token);
