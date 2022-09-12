SELECT 'CREATE DATABASE aims_enrollment_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'aims_enrollment_db')