#!/bin/sh

# backup the generated .sql
mkdir cmd/migrate/_migrations_back
mv cmd/migrate/migrations/00_initial.up.sql cmd/migrate/_migrations_back/00_initial.up.sql
# create new migration
go run cmd/gen/main.go -new
# apply new migration
go run cmd/migrate/main.go
