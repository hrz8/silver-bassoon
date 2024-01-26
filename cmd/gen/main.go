package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hrz8/silver-bassoon/pkg/gen"
	"github.com/hrz8/silver-bassoon/pkg/logger"

	_ "github.com/lib/pq"
)

//go:embed files/*.csv
var CsvFS embed.FS

var csvFiles = []string{
	"files/customer_companies.csv",
	"files/customers.csv",
	"files/orders.csv",
	"files/order_items.csv",
	"files/deliveries.csv",
}

var initialMigrationFile = "00_initial.up.sql"
var migrationFolderTarget = "cmd/migrate/migrations"

func main() {
	var newFile bool

	flag.BoolVar(&newFile, "new", false, "generate new migration file")

	flag.Parse()

	insert := insertTables(csvFiles)
	ddl := createTables(csvFiles, insert)

	writeToFile(ddl, newFile)
}

func createTables(files []string, insert string) string {
	var ddl string

	for _, file := range files {
		err := gen.ScanCreateTable(CsvFS, file, &ddl)
		if err != nil {
			panic(err)
		}
	}

	if gen.DDLCached != "" {
		return gen.DDLCached
	}

	sqlStatement, err := gen.AppropriateDDL(ddl, insert)
	if err != nil {
		return ddl
	}

	return sqlStatement
}

func insertTables(files []string) string {
	var insert string

	for _, file := range files {
		err := gen.ScanInsertTable(CsvFS, file, &insert)
		if err != nil {
			panic(err)
		}
	}

	return insert
}

func generateMigrationFilePrefix() string {
	currentTime := time.Now().UTC()

	prefix := currentTime.Format("20060102150405")

	return fmt.Sprintf("%s_initial.up.sql", prefix)
}

func writeToFile(sqlContent string, newFile bool) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	targetPath := strings.Split(migrationFolderTarget, "/")
	target := append([]string{cwd}, targetPath...)
	migrationsPath := filepath.Join(target...)

	filePath := filepath.Join(migrationsPath, initialMigrationFile)
	if newFile {
		filePath = filepath.Join(migrationsPath, generateMigrationFilePrefix())
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte(sqlContent), 0644); err != nil {
			log.Fatal(err)
		}

		logger.Info("new migration file created: %s\n", filePath)
	} else {
		logger.Info("migration file already exists: %s\n", filePath)
	}
}
