-- name: CreateRequest :exec
INSERT INTO requests (id, method, content, source_ip, response_code, timestamp, path)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: CreateHeaders :copyfrom
insert into request_headers (request_id, name, value)
values ($1, $2, $3);

-- name: CreateQueryParameters :copyfrom
insert into request_query_parameters (request_id, name, value)
values ($1, $2, $3);

-- name: GetRequests :many
select id, method, content, source_ip, response_code, timestamp, path
from requests
order by timestamp desc
limit $1;

-- name: GetRequestsPaged :many
select id, method, content, source_ip, response_code, timestamp, path
from requests
where id < $1
order by timestamp desc
limit $2;

-- name: GetHeadersById :many
select request_id, name, value
from request_headers
where request_id = $1;

-- name: GetQueryParamsById :many
select request_id, name, value
from request_query_parameters
where request_id = $1;

-- name: GetHeaders :many
select request_id, name, value
from request_headers;

-- name: GetQueryParams :many
select request_id, name, value
from request_query_parameters;

-- name: GetAllUsers :many
select *
from users;

-- name: GetUser :one
select *
from users
where username = $1;

-- name: UpdateUserPassword :exec
update users
set password = $1
where username = $2;

-- name: CreateUser :exec
insert into users (username, password)
values ($1, $2);

-- name: GetLatestKey :one
select *
from jwt_keys
order by id desc
limit 1;

-- name: CreateKey :exec
insert into jwt_keys (id, public_key, private_key)
values ($1, $2, $3);