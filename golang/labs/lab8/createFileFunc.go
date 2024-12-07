package lab8

import (
	"fmt"
	"os"
)

func CreateFile(filename string) (string, error) {
	_, err := os.Stat(filePath + filename)
	if err == nil {
		return "", fmt.Errorf("создание файла: файл %s уже существует", filename)
	}
	f, err := os.Create(filePath + filename)
	if err != nil {
		return "", fmt.Errorf("создание файла: %w", err)
	}
	defer f.Close()
	return filename, nil
}
