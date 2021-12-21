-- CREATE ROLE payment WITH LOGIN PASSWORD 'payment';
-- CREATE DATABASE payment_db;

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    iin text UNIQUE NOT NULL,
    login text UNIQUE NOT NULL,
    password text NOT NULL,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE DOMAIN role_name AS text CHECK (VALUE IN ('ROLE_ADMIN', 'ROLE_USER')); 

CREATE TABLE IF NOT EXISTS roles (
    id bigserial PRIMARY KEY,
    name role_name NOT NULL DEFAULT 'ROLE_USER',
    user_id integer UNIQUE REFERENCES users(id)
);