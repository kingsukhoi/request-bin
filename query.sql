-- name: CreateRequest :exec
INSERT INTO requests (id, method, content, source_ip, response_code, timestamp, path)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: CreateHeaders :copyfrom
insert into request_headers (request_id, name, value)
values ($1, $2, $3);

-- name: CreateQueryParameters :copyfrom
insert into request_query_parameters (request_id, name, value)
values ($1,$2,$3);

-- name: GetRequests :many
select id, method,content,source_ip,response_code,timestamp
from requests;

-- name: GetHeadersById :many
select request_id, name,value
from request_headers
where request_id = $1;

-- name: GetQueryParamsById :many
select request_id, name,value
from request_query_parameters
where request_id = $1;

-- name: GetHeaders :many
select request_id, name,value
from request_headers;

-- name: GetQueryParams :many
select request_id, name,value
from request_query_parameters;