CREATE DOMAIN role_name AS text CHECK (VALUE IN ('ROLE_ADMIN', 'ROLE_USER')); 

CREATE TABLE IF NOT EXISTS roles (
    id bigserial PRIMARY KEY,
    name role_name NOT NULL DEFAULT 'ROLE_USER',
    user_id integer UNIQUE REFERENCES users(id)
);