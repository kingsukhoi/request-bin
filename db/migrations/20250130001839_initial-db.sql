-- migrate:up
CREATE TABLE IF NOT EXISTS "requests"
(
    id       uuid primary key,
    method   text,
    content  bytea,
    source_ip text,
    response_code int
);
create table if not exists request_headers
(
    request_id uuid not null,
    name       text not null,
    value      text,
    foreign key (request_id) references requests (id)
);
create table if not exists request_query_parameters
(
    request_id uuid not null,
    name       text not null,
    value      text,
    foreign key (request_id) references requests (id)
);

-- migrate:down

drop table if exists request_query_parameters;
drop table if exists request_headers;
drop table if exists requests;