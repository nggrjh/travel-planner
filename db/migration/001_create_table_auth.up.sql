CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE auth (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    last_modified_at TIMESTAMP,
    CONSTRAINT uidx_auth_username UNIQUE (username),
    CONSTRAINT uidx_auth_email UNIQUE (email)
);
