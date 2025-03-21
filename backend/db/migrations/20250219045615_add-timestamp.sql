-- migrate:up
alter table requests
    add timestamp timestamptz default now();


-- migrate:down

alter table requests
    drop timestamp;