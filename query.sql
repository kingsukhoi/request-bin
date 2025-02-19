-- name: CreateRequest :exec
INSERT INTO requests (id, method, content, source_ip, response_code, timestamp)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: CreateHeaders :copyfrom
insert into request_headers (request_id, name, value)
values ($1, $2, $3);

-- name: CreateQueryParameters :copyfrom
insert into request_query_parameters (request_id, name, value)
values ($1,$2,$3);