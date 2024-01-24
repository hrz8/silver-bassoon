package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func connect(ctx context.Context) *pgx.Conn {
	connectionStr := "postgres://postgres:toor@localhost:5432/silver_bassoon?sslmode=disable"

	conn, err := pgx.Connect(ctx, connectionStr)
	if err != nil {
		panic(err)
	}

	return conn
}
