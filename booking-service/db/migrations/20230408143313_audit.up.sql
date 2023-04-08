create table public.audits
(
    id         bigserial primary key,
    name       text,
    action     varchar(50),
    identifier varchar(250),
    data       text,
    created_at timestamp with time zone default CURRENT_TIMESTAMP
);