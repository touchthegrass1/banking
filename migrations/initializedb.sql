CREATE DATABASE banking_db;
CREATE USER banking_db_user WITH PASSWORD 'examplepassword';
ALTER ROLE banking_db_user SET client_encoding TO 'utf8';
ALTER ROLE banking_db_user SET default_transaction_isolation TO 'serializable';
ALTER ROLE banking_db_user SET timezone TO 'UTC+4';
GRANT ALL PRIVILEGES ON DATABASE banking_db TO banking_db_user;

CREATE SCHEMA banking;
GRANT ALL ON SCHEMA banking TO banking_db_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA banking TO banking_db_user;
CREATE TYPE banking.client_type AS ENUM ('ie', 'individual', 'jp');

CREATE TYPE banking.card_type AS ENUM ('credit', 'debit');

CREATE TYPE banking.contract_type AS ENUM (
    'loan_agreement',
    'bank_account_agreement',
    'settlement_and_cash_service_agreement'
);

CREATE TABLE banking.client (
    client_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50),
    phone VARCHAR(12),
    registration_address VARCHAR(300),
    residential_address VARCHAR(300),
    client_type banking.client_type,
    ogrn int,
    inn int,
    kpp int
);

CREATE TABLE banking.card (
    card_id INT PRIMARY KEY,
    balance DECIMAL NOT NULL,
    valid_to TIMESTAMP,
    cvc_code VARCHAR(3),
    card_type banking.card_type,
    currency VARCHAR(3),
    client_id BIGINT REFERENCES banking.client(client_id) ON DELETE RESTRICT
);

CREATE TABLE banking.contract (
    contract_id BIGSERIAL PRIMARY KEY,
    contract_type banking.contract_type NOT NULL,
    conclusion_date DATE,
    contract_content TEXT,
    client_id BIGINT REFERENCES banking.client(client_id) ON DELETE RESTRICT
);

CREATE INDEX ON banking.contract (client_id);

CREATE TABLE banking.credit (
    credit_id BIGSERIAL PRIMARY KEY,
    summ DECIMAL NOT NULL,
    percent DECIMAL NOT NULL,
    conclusion_date DATE NOT NULL,
    end_date DATE NOT NULL,
    contract_id BIGINT UNIQUE REFERENCES banking.contract(contract_id) ON DELETE RESTRICT
);

CREATE TABLE banking.payment_schedule (
    payment_schedule_id BIGSERIAL PRIMARY KEY,
    total_summ DECIMAL NOT NULL,
    currency VARCHAR(3) NOT NULL,
    commision DECIMAL NOT NULL,
    repayment_of_interest_summ DECIMAL NOT NULL,
    summ_repayment_loan_part DECIMAL NOT NULL,
    date_begin DATE NOT NULL,
    date_end DATE NOT NULL,
    contract_id BIGINT REFERENCES banking.contract(contract_id) ON DELETE RESTRICT
);

CREATE INDEX ON banking.payment_schedule(contract_id);

ALTER ROLE banking_db_user SET search_path TO "$user", "public", "banking";
