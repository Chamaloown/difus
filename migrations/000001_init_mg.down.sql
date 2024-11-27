-- Drop tables in reverse order of creation to avoid foreign key constraints
DROP TABLE IF EXISTS almanax.users_jobs;
DROP TABLE IF EXISTS almanax.users;
DROP TABLE IF EXISTS almanax.jobs;
DROP TABLE IF EXISTS almanax.classes;
DROP TABLE IF EXISTS almanax.almanaxes;

-- Drop the schema
DROP SCHEMA IF EXISTS almanax;
