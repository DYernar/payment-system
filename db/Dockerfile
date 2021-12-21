FROM postgres:alpine
ADD scripts/000001_init_db.up.sql /docker-entrypoint-initdb.d
ADD scripts/000003_create_users_table.up.sql /docker-entrypoint-initdb.d
ADD scripts/000004_create_roles_table.up.sql /docker-entrypoint-initdb.d
EXPOSE 5432