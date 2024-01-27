package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func connect(ctx context.Context) *pgx.Conn {
	// postgres://postgres:toor@ip:5432/silver_bassoon?sslmode=disable
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return conn
}
