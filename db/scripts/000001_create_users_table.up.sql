CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    iin text UNIQUE NOT NULL,
    login text UNIQUE NOT NULL,
    password text NOT NULL,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);