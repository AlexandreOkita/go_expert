-- Escreve para banco postgres
create table products (
    id varchar(255) primary key,
    name varchar(255) not null,
    price numeric(10, 2) not null
);