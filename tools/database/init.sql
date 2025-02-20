CREATE ROLE task_manager WITH LOGIN PASSWORD '12345678';
ALTER ROLE task_manager CREATEDB;

-- Grant permissions only on the specific database
GRANT CONNECT ON DATABASE task_manager TO task_manager;
GRANT USAGE, SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO task_manager;
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA public TO task_manager;

-- Allow the user to have permissions on future tables and sequences
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO task_manager;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO task_manager;
