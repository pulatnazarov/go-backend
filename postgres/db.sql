CREATE TABLE countries (
    id SERIAL primary key,
    name varchar(30),
    currency varchar(10)
);

CREATE TABLE cities (
    id SERIAL primary key,
    name varchar(30),
    population int,
    country_id int references countries(id)
);

CREATE TABLE test (
    id SERIAL primary key,
    name varchar(30) not null,
    test text
);

create table users (
    id serial primary key,
    first_name varchar,
    last_name varchar,
    phone varchar
);

create table orders (
    id serial primary key,
    user_id int references users(id),
    total_sum int
);

create table products (
    id serial primary key,
    name varchar,
    quantity int,
    price int,
    order_id int references orders(id)
);

/*
table 1 - users (id, first_name, last_name, phone) 5
table 2 - orders (id, user_id, total_sum) 5
table 3 - products (id, name, quantity, price) 5
orders products relation

1. get all orders ( x <= total_sum <= y )
2. get all orders order by total_sum from low to high
3. get all users where first_name or last_name like 'some word'
4. get all products where quantity more than 5
*/