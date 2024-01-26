#!/bin/sh

rm -rf cmd/migrate/migrations
mkdir cmd/migrate/migrations
mv cmd/migrate/_migrations_back/00_initial.up.sql cmd/migrate/migrations/00_initial.up.sql
rm -rf cmd/migrate/_migrations_back
