-- Active: 1679718148710@@127.0.0.1@35432@account@public
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- best practice to use account role as enum
CREATE TYPE account_role AS ENUM ('admin', 'normal');
create table if not exists accounts(
id uuid primary key not null default uuid_generate_v4(),
username text not null,
password text not null,
role account_role not null,
created_at timestamptz not null default now(),
updated_at timestamptz not null default now(),
deleted_at timestamptz
);
CREATE INDEX accounts_deleted_at ON accounts(deleted_at);
CREATE UNIQUE INDEX accounts_unique_username ON accounts(username);

CREATE TYPE activity_type AS ENUM ('login', 'logout');
create table if not exists account_activities(
	id uuid primary key not null default uuid_generate_v4(),
	user_id uuid not null,
	type activity_type not null,
	created_at timestamptz not null default now(),
	updated_at timestamptz not null default now(),
	deleted_at timestamptz
);
CREATE INDEX accoount_activity_deleted_at ON account_activities(deleted_at);