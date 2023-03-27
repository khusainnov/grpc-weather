CREATE TABLE city_table
(
    id      smallserial  not null primary key,
    city    varchar(255) not null unique,
    counter int          not null
);