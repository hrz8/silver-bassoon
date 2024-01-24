package gen

import (
	"fmt"
	"strings"
)

func generateCreateTable(tableName string, header []string) string {
	columnDefs := make([]string, len(header))

	for i, columnName := range header {
		if columnName == "id" {
			columnDefs[i] = fmt.Sprintf("%s SERIAL PRIMARY KEY", columnName)
			continue
		}

		columnDefs[i] = fmt.Sprintf("%s VARCHAR(255)", columnName)
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(columnDefs, ", "))
}

func generateInsertTable(tableName string, cols []string) string {
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES ", tableName, strings.Join(cols, ", "))
}

func generateInsertValues(record []string, isLast bool) string {
	placeholder := "('%s'),"

	if isLast {
		placeholder = "('%s');"

	}

	return fmt.Sprintf(placeholder, strings.Join(record, "', '"))
}
