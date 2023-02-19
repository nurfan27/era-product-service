create table products
(
    id          serial primary key,
    name        text                    not null,
    price       integer                 not null,
    description text,
    quantity    integer   default 0     not null,
    created_at  timestamp default now() not null
);
