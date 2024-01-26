package main

import (
	"database/sql"
	"embed"
	"os"

	"github.com/hrz8/silver-bassoon/pkg/logger"
	"github.com/hrz8/silver-bassoon/pkg/migrator"
)

const migrationsDir = "migrations"

//go:embed migrations/*.sql
var MigrationsFS embed.FS

func main() {
	migrator := migrator.NewMigrator(MigrationsFS, migrationsDir)

	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = migrator.ApplyMigrations(conn)
	if err != nil {
		panic(err)
	}

	logger.Info("migrations done!")
}
