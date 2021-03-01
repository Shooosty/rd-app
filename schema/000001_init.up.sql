CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    role          varchar(255) not null,
    phone         varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE orders
(
    id            serial       not null unique,
    name          varchar(255) not null,
    user_id       varchar(255) not null
);
