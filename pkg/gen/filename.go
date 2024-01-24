package gen

import (
	"path"
	"strings"
)

func removeFileExtension(filename string, ext string) string {
	return strings.TrimSuffix(filename, ext)
}

func getFileName(filename string) string {
	return removeFileExtension(path.Base(filename), ".csv")
}
