CREATE USER "msTodo" WITH
    LOGIN
	NOSUPERUSER
	NOCREATEDB
	NOCREATEROLE
	NOINHERIT
	NOREPLICATION
	CONNECTION LIMIT -1
	PASSWORD 'xxxxxx';

CREATE DATABASE "msTodo"
    WITH 
    OWNER = "msTodo"
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

COMMENT ON DATABASE "msTodo"
    IS 'Don''t use default';

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(80) UNIQUE NOT NULL,
    image_uri VARCHAR NOT NULL,
    birthdate DATE NOT NULL,
    date_created TIMESTAMP NOT NULL    
);

CREATE TABLE lists(
    id SERIAL PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    description VARCHAR(80) UNIQUE NOT NULL,
    owner INT32 NOT NULL,
    date_created TIMESTAMP NOT NULL    
);

CREATE TABLE todos(
    id SERIAL PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    description VARCHAR(80) UNIQUE NOT NULL,
    owner INT32 NOT NULL,
    status INT32 NOT NULL,
    date_created TIMESTAMP NOT NULL  
    
);

CREATE TABLE statuses(

);

CREATE TABLE label(
);

CREATE TABLE comments(
);

CREATE TABLE reactions(
);
