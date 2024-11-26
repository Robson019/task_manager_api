create extension if not exists "uuid-ossp";

create table if not exists reset_password
(
    id uuid not null
        constraint df_reset_password_id default uuid_generate_v4()
        constraint pk_reset_password primary key,
    email varchar(256) not null,
    token text not null,
    created_at timestamp
        constraint df_reset_password_created_at default now()
);

create table if not exists role
(
    id uuid not null
        constraint df_role_id default uuid_generate_v4()
        constraint pk_role primary key,
    name varchar(200) not null
        constraint uk_role_name unique
);

create table if not exists account
(
    id uuid not null
        constraint df_account_id default uuid_generate_v4()
        constraint pk_account primary key,
    email varchar(256) not null
        constraint uk_account_email unique,
    password text not null,
    hash varchar(1024) not null,
    is_deleted bool not null
        constraint df_account_is_deleted default false,
    role_id uuid not null
        constraint fk_account_account_role_id references role (id)
);

create table if not exists task
(
    id uuid primary key default uuid_generate_v4(),
    title varchar(255) not null,
    description text,
    status varchar(10) not null check (status in ('pending', 'progress', 'done')),
    created_at timestamp default current_timestamp,
    updated_at timestamp default null,
    deleted_at timestamp default null
);
