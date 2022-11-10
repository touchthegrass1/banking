CREATE DATABASE banking_test_db;
CREATE USER banking_db_user WITH PASSWORD 'examplepassword';
ALTER ROLE banking_db_user SET client_encoding TO 'utf8';
ALTER ROLE banking_db_user SET default_transaction_isolation TO 'serializable';
ALTER ROLE banking_db_user SET timezone TO 'UTC+4';
GRANT ALL PRIVILEGES ON DATABASE banking_test_db TO banking_db_user;
