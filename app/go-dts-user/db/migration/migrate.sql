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

-- memiliki relasi dengan users table
-- dihubungkan dengan user_id
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

-- ON DELETE CASCADE -> kalau data di user kedelete
-- semua data yang berhubungan dengan user tersebut
-- di table user_photos akan kedelete juga

-- NOT NULL tidak boleh kosong
-- data di id TIDAK BOLEH NULL

-- PRIMARY KEY -> data unique di table tertentu
-- data ini tidak boleh ganda

-- sebenarnya kita bisa membuat kombinasi / 
-- data lain selain primary key
-- yang tidak boleh ganda dengan menggunakan -> INDEX
-- nik text UNIQUE, -- membuat unique index 
-- kenapa tidak NIK yang dijadikan PRIMARY KEY?
-- bisa bisa aja, tapi tidak common
-- id -> selain unique, dia harus generally berbeda

-- primary key:
-- serial
-- uuidV4
-- ulid

-- primary key di company A NIK 123456789 / id 123
-- primary key di company B NIK query?nik=123456789 / 123

-- Foreign Key -> penunjuk sebagai data yang diambil
-- dari table lain / penunjuk relasi dari table lain
-- untuk menghubungkan table
-- apakah foreign key wajib? 
-- FK menjadi wajib ketika kita ingin membuat table yang rapih / consistent
-- apakah akan berpengaruh ke query? tidak
-- tanpa FK, kita masih bisa menghubungkan table

-- serial -> int dengan characteristic auto increment
-- ketika kita tidak mendefine id, dia akan auto increment