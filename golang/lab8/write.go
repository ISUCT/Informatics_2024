package lab8

import (
	"fmt"
	"os"
)

func WriteFile(fileName string, newLine string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл")
	}
	defer file.Close()

	_, err = file.WriteString(newLine + "\n")
	if err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}
	return nil
}
