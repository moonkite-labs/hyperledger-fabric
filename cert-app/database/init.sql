CREATE DATABASE wallet;

CREATE TABLE IF NOT EXISTS Wallet (
    id serial PRIMARY KEY,
    label VARCHAR ( 50 ) UNIQUE NOT NULL,
    mspid VARCHAR ( 50 ) NOT NULL,
    public_key BYTEA UNIQUE NOT NULL,
    private_key BYTEA UNIQUE NOT NULL,
    created_at TIMESTAMP
);