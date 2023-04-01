-- Active: 1679718148710@@127.0.0.1@25432@user_management

-- DDL
CREATE DATABASE user_management;

CREATE TABLE users(  
    id serial NOT NULL PRIMARY KEY,
    email text UNIQUE,
    name VARCHAR(255),
    dob DATE,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
);

CREATE TABLE user_photos(  
    id serial NOT NULL PRIMARY KEY,
    url text,
    user_id int,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz,
    Foreign Key (user_id) 
        REFERENCES users(id)
);

CREATE TABLE user_photos_no_fk(  
    id serial NOT NULL PRIMARY KEY,
    url text,
    user_id int,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    deleted_at timestamptz
);

