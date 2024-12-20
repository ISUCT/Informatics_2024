package lab8

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(filepath string) (string, error) {
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения файла: %w", err)
	}
	return string(fileData), nil
}

func SearchInFile(filedata string, searchText string) bool {
	return strings.Contains(filedata, searchText)
}
