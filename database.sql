-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE estate (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	length INT NOT NULL,
	width INT NOT NULL
);

CREATE TABLE tree (
	estate_id UUID NOT NULL,
	x INT NOT NULL,
	y INT NOT NULL,
	height INT NOT NULL
);