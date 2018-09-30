SET TIME ZONE 'UTC';

DROP SCHEMA public CASCADE;

CREATE SCHEMA alpha;

-- CREATE ADMIN ROLE
CREATE ROLE admin WITH LOGIN SUPERUSER CREATEDB CREATEROLE REPLICATION NOINHERIT PASSWORD '<INSERT PASSWORD>';
GRANT ALL PRIVILEGES ON DATABASE grpcrud to admin;
ALTER ROLE admin SET search_path = alpha;

-- ALTER READ_WRITE USER ROLE
GRANT ALL PRIVILEGES ON DATABASE grpcrud to read_write_user;
GRANT CONNECT ON DATABASE grpcrud TO read_write_user;
GRANT USAGE ON SCHEMA alpha TO read_write_user;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA alpha TO read_write_user;
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA alpha TO read_write_user;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA alpha TO read_write_user;
ALTER ROLE read_write_user SET search_path = alpha;
ALTER ROLE read_write_user WITH NOINHERIT;

create table if not exists alpha.todo (
    todo_id                 serial not null,
    title                   varchar(200) default null,
    description             varchar(1024) default null,
    reminder                timestamptz default null,
    constraint todo_pkey    primary key (todo_id)
);
