package lab8

import (
	"fmt"
	"os"
)

func CreatingFile(fileName string) error {
	_, err := os.Stat(fileName)
	if err == nil {
		return fmt.Errorf("файл с таким именем уже создан")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("произошла ошибка при создании файла: %w", err)
	}
	defer file.Close()

	return nil
}
