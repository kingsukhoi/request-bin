-- migrate:up
CREATE TABLE users
(
    username text primary key,
    password text not null
);
CREATE TABLE jwt_keys
(
    id          uuid primary key,
    public_key  text not null,
    private_key text not null
);

-- migrate:down

drop table jwt_keys;
drop table users;