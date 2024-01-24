package csv

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io/fs"
	"log"
	"strings"
)

func getReader(files embed.FS, fileName string) (fs.File, *csv.Reader, error) {
	file, err := files.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open file: %v", err)
	}

	reader := csv.NewReader(file)

	return file, reader, nil
}

func ReadAll(files embed.FS, fileName string, csvString *string) error {
	file, reader, err := getReader(files, fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		*csvString += strings.Join(line, ",") + "\n"
	}

	return nil
}
