CREATE DATABASE crud_app;

\c crud_app

CREATE TABLE items (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
description TEXT,
price NUMERIC(10, 2)
);