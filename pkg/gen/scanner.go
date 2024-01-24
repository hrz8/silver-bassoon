package gen

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
)

func ScanCreateTable(files embed.FS, fileName string, ddl *string) error {
	tableName := getFileName(fileName)

	file, err := files.Open(fileName)
	if err != nil {
		return fmt.Errorf("unable to open csv file: %v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("unable to read csv file: %v", err)
	}

	sqlStatement := generateCreateTable(tableName, header)
	*ddl += sqlStatement

	return nil
}

func ScanInsertTableLines(files embed.FS, fileName string) (int, error) {
	file, err := files.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("unable to open csv file: %v", err)
	}

	defer file.Close()

	allReader := csv.NewReader(file)

	lines, err := allReader.ReadAll()
	if err != nil {
		return 0, fmt.Errorf("unable to read csv file: %v", err)
	}

	return len(lines), nil
}

func ScanInsertTable(files embed.FS, fileName string, insert *string) error {
	lineLength, err := ScanInsertTableLines(files, fileName)
	if err != nil {
		return fmt.Errorf("unable count csv lines: %v", err)
	}

	tableName := getFileName(fileName)

	file, err := files.Open(fileName)
	if err != nil {
		return fmt.Errorf("unable to open csv file: %v", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("unable to read csv file: %v", err)
	}

	sqlStatement := generateInsertTable(tableName, header)

	currentLine := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("unable to detailing csv file: %v", err)
		}

		currentLine++
		sqlStatement += generateInsertValues(record, currentLine == lineLength)
	}

	*insert += sqlStatement

	return nil
}
