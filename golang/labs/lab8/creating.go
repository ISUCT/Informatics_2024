package lab8

import (
	"fmt"
	"os"
)

func CreateFile(fileName string) error {
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("файл с таким именем уже существует")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer file.Close()
	return nil
}
