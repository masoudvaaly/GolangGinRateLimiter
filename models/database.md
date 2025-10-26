CREATE DATABASE crud_app;

\c crud_app

CREATE TABLE items (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
description TEXT,
price NUMERIC(10, 2)
);

CREATE TABLE nav (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
description TEXT,
calcDate TIMESTAMP,
light BOOLEAN DEFAULT 1
);

CREATE TABLE FiscalYear (
id SERIAL PRIMARY KEY,
year_name TEXT NOT NULL,
is_active BOOLEAN DEFAULT 1
start_date TIMESTAMP,
end_date TIMESTAMP
);
