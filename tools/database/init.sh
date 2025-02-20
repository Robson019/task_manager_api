#!/bin/sh
set -e

# Cria a role task_manager, utilizando a variável de ambiente POSTGRES_PASSWORD
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE ROLE task_manager WITH LOGIN PASSWORD '${POSTGRES_PASSWORD}';
ALTER ROLE task_manager CREATEDB;
EOSQL

# Executa os comandos de permissão
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-'EOSQL'
    -- Concede permissões somente na base de dados específica
    GRANT CONNECT ON DATABASE task_manager TO task_manager;
    GRANT USAGE, SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO task_manager;
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA public TO task_manager;

-- Permite que o usuário tenha permissões em tabelas e sequências futuras
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO task_manager;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO task_manager;
EOSQL
