FROM postgres:alpine

ADD /postgres/migrations/01_schema.sql /docker-entrypoint-initdb.d
ADD /postgres/migrations/02_seed_data.sql /docker-entrypoint-initdb.d

RUN chmod a+r /docker-entrypoint-initdb.d/*
