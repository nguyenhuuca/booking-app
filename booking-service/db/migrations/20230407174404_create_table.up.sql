create table public.branches
(
    id bigserial primary key,
    name       text,
    created_at timestamp with time zone default CURRENT_TIMESTAMP
);


create table public.products
(
    id         bigserial primary key,
    name       text,
    branch_id  bigint,
    price      numeric(12,2),
    created_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table public.products
    owner to postgres;




