CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    last_modified_at TIMESTAMP,
    CONSTRAINT uidx_users_email UNIQUE (email)
);
