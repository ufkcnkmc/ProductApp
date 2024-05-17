#!/bin/bash

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest

echo "Postgresql starting"
sleep 3

docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE productapp"
sleep 3
echo "Database productapp created"

docker exec -it postgres-test psql -U postgres -d productapp -c "
CREATE TABLE IF NOT EXISTS products (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  price DOUBLE PRECISION NOT NULL,
  discount DOUBLE PRECISION,
  store VARCHAR(255) NOT NULL
);
"
sleep 3
echo "Table products created"
