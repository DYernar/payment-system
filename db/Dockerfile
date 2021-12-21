FROM postgres:alpine
ADD scripts/000001_create_users_table.up.sql /docker-entrypoint-initdb.d
ADD scripts/000002_create_roles_table.up.sql /docker-entrypoint-initdb.d
EXPOSE 5432