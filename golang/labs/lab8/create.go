package lab8

import (
	"fmt"
	"io"
	"os"
)

func CreateFile(filename string) (string, error) {
	if _, err := os.Stat(filename); err == nil {
		return "", fmt.Errorf("Файл %s уже существует", filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("Ошибка создания файла: %w", err)
	}
	defer file.Close()
	return filename, nil
}

func WriteToFile(filename string, info string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла: %w", err)
	}
	defer file.Close()
	_, err = io.WriteString(file, info)
	if err != nil {
		return fmt.Errorf("Ошибка при записи в файл: %w", err)
	}
	return nil
}
