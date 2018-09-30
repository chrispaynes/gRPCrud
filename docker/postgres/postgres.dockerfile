FROM postgres:10.5
ADD ./docker/postgres/init.sql /docker-entrypoint-initdb.d/