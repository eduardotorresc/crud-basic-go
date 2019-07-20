create database test;

use test;

create table product(
    id int auto_increment primary key,
    description varchar(255),
    quantify int,
    price decimal(10,2),
    amount decimal(10,2),
    created_at timestamp default current_timestamp()
);

insert into product (description, quantify, price) values('Orange', 50, 2.50);
